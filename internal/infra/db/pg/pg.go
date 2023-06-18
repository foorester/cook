package pg

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/foorester/cook/internal/sys"

	"github.com/foorester/cook/internal/sys/errors"
)

type (
	DB struct {
		sys.Core
		db *sql.DB
	}
)

const (
	name = "pg-db"
)

func NewDB(opts ...sys.Option) *DB {
	return &DB{
		Core: sys.NewCore(name, opts...),
		db:   nil,
	}
}

func (db *DB) Start(ctx context.Context) error {
	return db.Connect()
}

func (db *DB) Connect() error {
	pgdb, err := sql.Open("pgx", db.connString())
	if err != nil {
		msg := fmt.Sprintf("%s connection error", db.Name())
		return errors.Wrap(msg, err)
	}

	err = pgdb.Ping()
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", db.Name())
		return errors.Wrap(msg, err)
	}

	db.Log().Infof("%s database connected!", db.Name())
	return nil
}

func (db *DB) DB() any {
	return db.db
}

func (db *DB) connString() (connString string) {
	cfg := db.Cfg()
	user := cfg.GetString("db.pg.user")
	pass := cfg.GetString("db.pg.pass")
	name := cfg.GetString("db.pg.db")
	host := cfg.GetString("db.pg.host")
	port := cfg.GetInt("db.pg.port")
	schema := cfg.GetString("db.pg.schema")
	sslMode := cfg.GetBool("db.pg.sslmode")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d search_path=%s", user, pass, name, host, port, schema)

	if sslMode {
		connStr = connStr + " sslmode=disable"
	} else {
		connStr = connStr + " sslmode=disable"
	}

	return connStr
}
