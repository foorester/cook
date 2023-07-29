package sqlite

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/infra/seed"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/errors"
)

const (
	seedTable = "seeds"
	seedPath  = "assets/seeding/sqlite"
)

type (
	// Seeder struct.
	Seeder struct {
		sys.Core
		assetsPath string
		fs         embed.FS
		db         *sql.DB
		steps      []Seed
	}

	// Seed struct
	Seed struct {
		Order    int
		Executor seed.Exec
	}

	seedRecord struct {
		ID        uuid.UUID      `dbPath:"id" json:"id"`
		Index     sql.NullInt64  `dbPath:"index" json:"index"`
		Name      sql.NullString `dbPath:"name" json:"name"`
		CreatedAt db.NullTime    `dbPath:"created_at" json:"createdAt"`
	}
)

type (
	SeedRecord struct {
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

func NewSeeder(fs embed.FS, opts ...sys.Option) (mig *Seeder) {
	m := &Seeder{
		Core:       sys.NewCore("seeder", opts...),
		assetsPath: seedPath,
		fs:         fs,
	}

	return m
}

func (s *Seeder) SetAssetsPath(path string) {
	s.assetsPath = path
}

func (s *Seeder) AssetsPath() string {
	return s.assetsPath
}

func (s *Seeder) Start(ctx context.Context) error {
	s.Log().Infof("%s started", s.Name())

	err := s.Connect()
	if err != nil {
		return errors.Wrapf(err, "%s start error", s.Name())
	}

	err = s.addSteps()
	if err != nil {
		return errors.Wrapf(err, "%s start error", s.Name())
	}

	return s.Seed()
}

func (s *Seeder) Connect() error {
	path := s.Cfg().GetString(config.Key.SQLiteFilePath)
	sqlDB, err := sql.Open("sqlite3", path+"?_journal_mode=WAL")

	if err != nil {
		return errors.Wrapf(err, "%s connection error", s.Name())
	}

	err = sqlDB.Ping()
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", s.Name())
		return errors.Wrap(err, msg)
	}

	s.Log().Infof("%s database connected", s.Name())

	s.db = sqlDB

	return nil
}

// GetTx returns a new transaction from seeder connection
func (m *Seeder) GetTx() (tx *sql.Tx, err error) {
	tx, err = m.db.Begin()
	if err != nil {
		return tx, err
	}

	return tx, nil
}

// PreSetup creates database
// and seeds table if needed.
func (s *Seeder) PreSetup() (err error) {
	if !s.seedsTableExists() {
		err = s.createSeedsTable()
		if err != nil {
			return err
		}
	}

	return nil
}

// dbExists returns true if migrator referenced database has been already created.
func (s *Seeder) dbExists() bool {
	st := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='database' AND name='%s';", s.Name())

	rows, err := s.db.Query(st)
	if err != nil {
		s.Log().Infof("Error checking database: %w", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var dbName string
		err = rows.Scan(&dbName)
		if err != nil {
			s.Log().Errorf("Cannot read query result: %w", err)
			return false
		}
		return true
	}

	return false
}

// seedsTableExists returns true if seed table exists.
func (s *Seeder) seedsTableExists() bool {
	st := fmt.Sprintf("SELECT name FROM sqlite_master WHERE type='table' AND name='%s';", seedTable)

	rows, err := s.db.Query(st)
	if err != nil {
		s.Log().Errorf("Error checking database: %s", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var tableName string
		err = rows.Scan(&tableName)
		if err != nil {
			s.Log().Errorf("Cannot read query result: %s\n", err)
			return false
		}

		return true
	}

	return false
}

func (s *Seeder) CloseAppConns() (string, error) {
	dbName := s.Cfg().GetString(config.Key.SQLiteFilePath)

	err := s.db.Close()
	if err != nil {
		return dbName, err
	}

	adminConn, err := sql.Open("sqlite3", s.Name())
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

func (s *Seeder) createSeedsTable() (err error) {
	tx, err := s.GetTx()
	if err != nil {
		return err
	}

	query := `CREATE TABLE %s (
    id UUID PRIMARY KEY,
    idx INTEGER,
    name VARCHAR(64),
    created_at TEXT
    );`

	st := fmt.Sprintf(query, seedTable)

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

func (s *Seeder) AddSeed(o int, e seed.Exec) {
	mig := Seed{Order: o, Executor: e}
	s.steps = append(s.steps, mig)
}

func (s *Seeder) Seed() (err error) {
	err = s.PreSetup()
	if err != nil {
		return errors.Wrap(err, "seeding error")
	}

	for i, _ := range s.steps {
		mg := s.steps[i]
		exec := mg.Executor
		idx := exec.GetIndex()
		name := exec.GetName()
		seedFxs := exec.GetSeeds()

		// Get a new Tx from migrator
		tx, err := s.GetTx()
		if err != nil {
			return errors.Wrap(err, "seeding error")
		}

		//Continue if already applied
		if !s.canApplySeed(idx, name, tx) {
			s.Log().Infof("Seed '%s' already applied", name)
			tx.Commit() // No need to handle eventual error here
			continue
		}

		for _, sfx := range seedFxs {
			err = sfx(tx)
			if err != nil {
				break
			}
		}

		if err != nil {
			s.Log().Infof("%s seed not executed", name)
			err2 := tx.Rollback()
			if err2 != nil {
				return errors.Wrap(err2, "seeding rollback error")
			}

			return errors.Wrapf(err, "cannot run seed '%s'", name)
		}

		// Register seed
		exec.SetTx(tx)
		err = s.recSeed(exec)

		err = tx.Commit()
		if err != nil {
			msg := fmt.Sprintf("Cannot update seed table: %s\n", err.Error())
			s.Log().Errorf("seeding commit error: %s", msg)
			err = tx.Rollback()
			if err != nil {
				return errors.Wrap(err, "seeding rollback error")
			}
			return errors.New(msg)
		}

		s.Log().Infof("Seed executed: %s", name)
	}

	return nil
}

func (s *Seeder) recSeed(e seed.Exec) error {
	query := `INSERT INTO %s (id, idx, n VALUES (:id, :idx, :name, :created_at);`

	st := fmt.Sprintf(query, seedTable)

	uid, err := uuid.NewUUID()
	if err != nil {
		return errors.Wrap(err, "cannot update seed table")
	}

	_, err = e.GetTx().Exec(st,
		ToNullString(uid.String()),
		ToNullInt64(e.GetIndex()),
		ToNullString(e.GetName()),
		ToNullString(time.Now().Format(time.RFC3339)),
	)

	if err != nil {
		return errors.Wrap(err, "cannot update seed table")
	}

	return nil
}

func (s *Seeder) cancelRollback(index int64, name string, tx *sql.Tx) bool {
	query := `SELECT (COUNT(*) > 0) AS record_exists FROM %s 
                                       WHERE idx = %d 
                                           :AND name = '%s'`

	st := fmt.Sprintf(query, seedTable, index, name)
	r, err := tx.Query(st)

	if err != nil {
		s.Log().Errorf("Cannot determine rollback status: %w", err)
		return true
	}

	for r.Next() {
		var applied sql.NullBool
		err = r.Scan(&applied)
		if err != nil {
			s.Log().Errorf("Cannot determine seed status: %w", err)
			return true
		}

		return !applied.Bool
	}

	return true
}

func (s *Seeder) canApplySeed(index int64, name string, tx *sql.Tx) bool {
	query := `SELECT (COUNT(*) > 0) AS record_exists FROM %s 
                                       WHERE idx = %d 
                                           AND name = '%s'`

	st := fmt.Sprintf(query, seedTable, index, name)
	r, err := tx.Query(st)
	defer r.Close()

	if err != nil {
		s.Log().Errorf("Cannot determine seed status: %w", err)
		return false
	}

	for r.Next() {
		var exists sql.NullBool
		err = r.Scan(&exists)
		if err != nil {
			s.Log().Errorf("Cannot determine seed status: %s", err)
			return false
		}

		return !exists.Bool
	}

	return true
}

func (s *Seeder) addSteps() error {
	qq, err := s.readInsertSets()
	if err != nil {
		return err
	}

	for i, q := range qq {
		var seeds []seed.SeedFx
		for _, i := range q.Inserts {
			seeds = append(seeds, s.genTxExecFunc(i))
		}

		step := &step{
			Index: q.Index,
			Name:  q.Name,
			Seeds: seeds,
		}

		s.AddSeed(i, step)
	}

	return nil
}

func (s *Seeder) genTxExecFunc(query string) func(tx *sql.Tx) error {
	return func(tx *sql.Tx) error {
		_, err := tx.Exec(query)
		return err
	}
}

type insertSet struct {
	Index   int64
	Name    string
	Inserts []string
}

func (s *Seeder) readInsertSets() ([]insertSet, error) {
	var iiss []insertSet

	files, err := s.fs.ReadDir(s.assetsPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}

		filePath := fmt.Sprintf("%s/%s", s.assetsPath, file.Name())
		content, err := s.fs.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		var statements []string
		insertStmts := strings.Split(string(content), "--SEED")
		if len(insertStmts) < 1 {
			msg := fmt.Sprintf("invalid seed file format: %s", file.Name())
			return nil, errors.New(msg)
		}

		for _, istmt := range insertStmts {
			insertSt := strings.TrimSpace(strings.TrimPrefix(istmt, "--SEED\n"))
			statements = append(statements, insertSt)
		}

		idx, name := stepName(filePath)

		is := insertSet{
			Index:   idx,
			Name:    name,
			Inserts: statements,
		}

		iiss = append(iiss, is)
	}

	return iiss, nil
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

func (s *Seeder) count() (last int) {
	return len(s.steps)
}

func (s *Seeder) last() (last int) {
	return s.count() - 1
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
