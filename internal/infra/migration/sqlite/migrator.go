package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/infra/migration"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/errors"
)

const (
	migTable = "migrations"
	migPath  = "assets/migrations/sqlite"
)

type (
	// MigFx type alias

	// Migrator struct.
	Migrator struct {
		sys.Core
		assetsPath string
		fs         embed.FS
		db         *sql.DB
		steps      []Migration
	}

	// Exec interface.

	// Migration struct.
	Migration struct {
		Order    int
		Executor migration.Exec
	}

	migRecord struct {
		ID        uuid.UUID      `dbPath:"id" json:"id"`
		Index     sql.NullInt64  `dbPath:"index" json:"index"`
		Name      sql.NullString `dbPath:"name" json:"name"`
		CreatedAt db.NullTime    `dbPath:"created_at" json:"createdAt"`
	}
)

type (
	MigrationRecord struct {
		ID        string
		Index     int
		Name      string
		CreatedAt string
	}
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func NewMigrator(fs embed.FS, opts ...sys.Option) (mig *Migrator) {
	m := &Migrator{
		Core:       sys.NewCore("migrator", opts...),
		assetsPath: migPath,
		fs:         fs,
	}

	return m
}

func (m *Migrator) SetAssetsPath(path string) {
	m.assetsPath = path
}

func (m *Migrator) AssetsPath() string {
	return m.assetsPath
}

func (m *Migrator) Start(ctx context.Context) error {
	m.Log().Infof("%s started", m.Name())

	err := m.Connect()
	if err != nil {
		return errors.Wrapf(err, "%s start error", m.Name())
	}

	err = m.addSteps()
	if err != nil {
		return errors.Wrapf(err, "%s start error", m.Name())
	}

	return m.Migrate()
}

func (m *Migrator) Connect() error {
	path := m.Cfg().GetString(config.Key.SQLiteFilePath)

	sqlDB, err := sql.Open("sqlite3", path+"?_journal_mode=WAL")
	if err != nil {
		return errors.Wrapf(err, "%s connection error", m.Name())
	}

	err = sqlDB.Ping()
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", m.Name())
		return errors.Wrap(err, msg)
	}

	m.Log().Infof("%s database connected", m.Name())

	m.db = sqlDB

	return nil
}

// GetTx returns a new transaction from migrator connection
func (m *Migrator) GetTx() (tx *sql.Tx, err error) {
	tx, err = m.db.Begin()
	if err != nil {
		return tx, err
	}

	return tx, nil
}

func (m *Migrator) PreSetup() (err error) {
	if !m.migTableExists() {
		err = m.createMigrationsTable()
		if err != nil {
			return err
		}
	}

	return nil
}

// dbExists returns true if migrator referenced database has been already created.
func (m *Migrator) dbExists() bool {
	st := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='database' AND name='%s';", m.Name())

	rows, err := m.db.Query(st)
	if err != nil {
		m.Log().Infof("Error checking database: %w", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		if err != nil {
			m.Log().Errorf("Cannot read query result: %w", err)
			return false
		}
		return true
	}

	return false
}

// migTableExists returns true if migration table exists.
func (m *Migrator) migTableExists() bool {
	st := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", migTable)

	rows, err := m.db.Query(st)
	if err != nil {
		m.Log().Errorf("Error checking database: %s", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			m.Log().Errorf("Cannot read query result: %s\n", err)
			return false
		}

		return true
	}

	return false
}

// DropDB migration.
func (m *Migrator) DropDB() (dbPath string, err error) {
	dbPath, err = m.CloseAppConns()
	if err != nil {
		return dbPath, errors.Wrap(err, "drop db error")
	}

	// Close the SQLite connection before dropping the database file
	err = m.db.Close()
	if err != nil {
		m.Log().Errorf("drop dbPath error: %w", err) // Maybe it was already closed.
	}

	err = os.Remove(dbPath)
	if err != nil {
		return dbPath, err
	}

	return dbPath, nil
}

func (m *Migrator) CloseAppConns() (string, error) {
	dbName := m.Cfg().GetString(config.Key.SQLiteFilePath)

	err := m.db.Close()
	if err != nil {
		return dbName, err
	}

	adminConn, err := sql.Open("sqlite3", m.Name())
	if err != nil {
		return dbName, err
	}
	defer adminConn.Close()

	// Terminate all connections to the database (SQLite does not support concurrent connections)
	st := fmt.Sprintf(`PRAGMA busy_timeout = 5000;`)
	_, err = adminConn.Exec(st)
	if err != nil {
		return dbName, err
	}

	return dbName, nil
}

func (m *Migrator) createMigrationsTable() (err error) {
	tx, err := m.GetTx()
	if err != nil {
		return err
	}

	query := `CREATE TABLE %s (
    id UUID PRIMARY KEY,
    idx INTEGER,
    name VARCHAR(64),
    created_at TEXT
    );`

	st := fmt.Sprintf(query, migTable)

	_, err = tx.Exec(st)
	if err != nil {
		err2 := tx.Rollback()
		if err2 != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}

func (m *Migrator) AddMigration(o int, e migration.Exec) {
	mig := Migration{Order: o, Executor: e}
	m.steps = append(m.steps, mig)
}

func (m *Migrator) Migrate() (err error) {
	err = m.PreSetup()
	if err != nil {
		return errors.Wrap(err, "migrate error")
	}

	for i, _ := range m.steps {
		mg := m.steps[i]
		exec := mg.Executor
		idx := exec.GetIndex()
		name := exec.GetName()
		upFx := exec.GetUp()

		// Get a new Tx from migrator
		tx, err := m.GetTx()
		if err != nil {
			return errors.Wrap(err, "migrate error")
		}

		//Continue if already applied
		if !m.canApplyMigration(idx, name, tx) {
			m.Log().Infof("Migration '%s' already applied", name)
			tx.Commit() // No need to handle eventual error here
			continue
		}

		err = upFx(tx)

		if err != nil {
			m.Log().Infof("%s migration not executed", name)
			err2 := tx.Rollback()
			if err2 != nil {
				return errors.Wrap(err2, "migrate rollback error")
			}

			return errors.Wrapf(err, "cannot run migration '%s'", name)
		}

		// Register migration
		exec.SetTx(tx)
		err = m.recMigration(exec)

		err = tx.Commit()
		if err != nil {
			msg := fmt.Sprintf("Cannot update migration table: %s\n", err.Error())
			m.Log().Errorf("migrate commit error: %s", msg)
			err = tx.Rollback()
			if err != nil {
				return errors.Wrap(err, "migrate rollback error")
			}
			return errors.New(msg)
		}

		m.Log().Infof("Migration executed: %s", name)
	}

	return nil
}

// Rollback migration.
func (m *Migrator) Rollback(steps ...int) error {
	// Default to 1 step if no value is provided
	s := 1
	if len(steps) > 0 && steps[0] > 1 {
		s = steps[0]
	}

	// Default to max n° migration if steps is higher
	c := m.count()
	if s > c {
		s = c
	}

	m.rollback(s)
	return nil
}

// RollbackAll migration.
func (m *Migrator) RollbackAll() error {
	return m.rollback(m.count())
}

func (m *Migrator) rollback(steps int) error {
	processed := 0
	count := m.count()

	for i := count - 1; i >= 0; i-- {
		mg := m.steps[i]
		exec := mg.Executor
		idx := exec.GetIndex()
		name := exec.GetName()
		downFx := exec.GetDown()

		// Get a new Tx from migrator
		tx, err := m.GetTx()
		if err != nil {
			return errors.Wrap(err, "rollback error")
		}

		// Continue if already applied
		if !m.canApplyRollback(idx, name, tx) {
			m.Log().Infof("Rollback '%s' cannot be executed", name)
			tx.Commit() // No need to handle eventual error here
			continue
		}

		// Pass Tx to the executor
		err = downFx(tx)
		if err != nil {
			m.Log().Infof("%s rollback not executed", name)
			err2 := tx.Rollback()
			if err2 != nil {
				return errors.Wrap(err2, "rollback rollback error")
			}
			return errors.Wrapf(err, "cannot run rollback '%s'", name)
		}

		// Remove migration record
		exec.SetTx(tx)
		err = m.delMigration(exec)

		err = tx.Commit()
		if err != nil {
			msg := fmt.Sprintf("Cannot delete migration table: %s\n", err.Error())
			m.Log().Errorf("rollback commit error: %s", msg)
			err = tx.Rollback()
			if err != nil {
				return errors.Wrap(err, "rollback rollback error")
			}
			return errors.New(msg)
		}

		processed++
		if processed == steps {
			m.Log().Infof("Rollback executed: %s", name)
			return nil
		}
	}

	return nil
}

func (m *Migrator) SoftReset() error {
	err := m.RollbackAll()
	if err != nil {
		log.Printf("Cannot rollback database: %s", err.Error())
		return err
	}

	err = m.Migrate()
	if err != nil {
		log.Printf("Cannot migrate database: %s", err.Error())
		return err
	}

	return nil
}

func (m *Migrator) Reset() error {
	_, err := m.DropDB()
	if err != nil {
		m.Log().Errorf("Drop database error: %w", err)
		// Don't return maybe it was not created before.
	}

	err = m.Migrate()
	if err != nil {
		return errors.Wrap(err, "drop database error")
	}

	return nil
}

func (m *Migrator) recMigration(e migration.Exec) error {
	query := `INSERT INTO %s (id, idx, name, created_at) VALUES (:id, :idx, :name, :created_at);`

	st := fmt.Sprintf(query, migTable)

	uid, err := uuid.NewUUID()
	if err != nil {
		return errors.Wrap(err, "cannot update migration table")
	}

	_, err = e.GetTx().Exec(st,
		ToNullString(uid.String()),
		ToNullInt64(e.GetIndex()),
		ToNullString(e.GetName()),
		ToNullString(time.Now().Format(time.RFC3339)),
	)

	if err != nil {
		m.Log().Error(err)
		return errors.Wrap(err, "cannot update migration table")
	}

	return nil
}

func (m *Migrator) cancelRollback(index int64, name string, tx *sql.Tx) bool {
	query := `SELECT (COUNT(*) > 0) AS record_exists 
		FROM %s 
		WHERE idx = %d 
		    AND name = '%s'`
	st := fmt.Sprintf(query, migTable, index, name)
	r, err := tx.Query(st)

	if err != nil {
		m.Log().Errorf("Cannot determine rollback status: %w", err)
		return true
	}

	for r.Next() {
		var applied sql.NullBool
		err = r.Scan(&applied)
		if err != nil {
			m.Log().Errorf("Cannot determine migration status: %w", err)
			return true
		}

		return !applied.Bool
	}

	return true
}

func (m *Migrator) canApplyMigration(index int64, name string, tx *sql.Tx) bool {
	query := `SELECT (COUNT(*) > 0) AS record_exists 
		FROM %s 
		WHERE idx = %d 
		    AND name = '%s'`

	st := fmt.Sprintf(query, migTable, index, name)
	r, err := tx.Query(st)
	defer r.Close()

	if err != nil {
		m.Log().Errorf("Cannot determine migration status: %w", err)
		return false
	}

	for r.Next() {
		var exists sql.NullBool
		err = r.Scan(&exists)
		if err != nil {
			m.Log().Errorf("Cannot determine migration status: %s", err)
			return false
		}

		return !exists.Bool
	}

	return true
}

func (m *Migrator) canApplyRollback(index int64, name string, tx *sql.Tx) bool {
	return !m.canApplyMigration(index, name, tx)
}

func (m *Migrator) delMigration(e migration.Exec) error {
	idx := e.GetIndex()
	name := e.GetName()

	query := `DELETE FROM %s WHERE idx = %d AND name = '%s'`

	st := fmt.Sprintf(query, migTable, idx, name)
	_, err := e.GetTx().Exec(st)

	if err != nil {
		return errors.Wrap(err, "cannot delete migration table record")
	}

	return nil
}

func (m *Migrator) addSteps() error {
	qq, err := m.readMigQueries()
	if err != nil {
		return err
	}

	for i, q := range qq {
		s := &step{
			Index: q.Index,
			Name:  q.Name,
			Up:    m.genTxExecFunc(q.Up),
			Down:  m.genTxExecFunc(q.Down),
		}

		m.AddMigration(i, s)
	}

	return nil
}

func (m *Migrator) genTxExecFunc(query string) func(tx *sql.Tx) error {
	return func(tx *sql.Tx) error {
		_, err := tx.Exec(query)
		return err
	}
}

type queries struct {
	Index int64
	Name  string
	Up    string
	Down  string
}

func (m *Migrator) readMigQueries() ([]queries, error) {
	var qq []queries

	files, err := m.fs.ReadDir(m.assetsPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", m.assetsPath, file.Name())
		content, err := m.fs.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		sections := strings.Split(string(content), "--DOWN")
		if len(sections) < 2 {
			msg := fmt.Sprintf("invalid migration file format: %s", file.Name())
			return nil, errors.New(msg)
		}

		up := strings.TrimSpace(strings.TrimPrefix(sections[0], "--UP\n"))
		down := strings.TrimSpace(sections[1])

		idx, name := stepName(filePath)

		q := queries{
			Index: idx,
			Name:  name,
			Up:    up,
			Down:  down,
		}

		qq = append(qq, q)
	}

	return qq, nil
}

func stepName(filename string) (idx int64, name string) {
	base := filepath.Base(filename)
	base = strings.TrimSuffix(base, filepath.Ext(base))

	re := regexp.MustCompile(`^[-\d]+`)
	indexStr := re.FindString(base)
	idx, _ = strconv.ParseInt(strings.TrimSuffix(indexStr, "-"), 10, 64)

	name = re.ReplaceAllString(base, "")
	name = strings.ReplaceAll(name, "-", " ")
	name = strings.ToLower(name)
	name = strings.ReplaceAll(name, " ", "-")

	return idx, name
}

func (m *Migrator) count() (last int) {
	return len(m.steps)
}

func (m *Migrator) last() (last int) {
	return m.count() - 1
}

func getFxName(i interface{}) string {
	n := runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	t := strings.FieldsFunc(n, split)
	return t[len(t)-2]
}

func split(r rune) bool {
	return r == '.' || r == '-'
}

func migName(upFxName string) string {
	return toSnakeCase(upFxName)
}

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func ToNullTime(t time.Time) db.NullTime {
	return db.NullTime{
		Time:  t,
		Valid: true,
	}
}

func ToNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}

func ToNullInt(i int64) sql.NullInt32 {
	return sql.NullInt32{
		Int32: int32(i),
		Valid: true,
	}
}

func ToNullInt64(i int64) sql.NullInt64 {
	return sql.NullInt64{
		Int64: i,
		Valid: true,
	}
}

func ToNullBool(b bool) sql.NullBool {
	return sql.NullBool{
		Bool:  b,
		Valid: true,
	}
}
