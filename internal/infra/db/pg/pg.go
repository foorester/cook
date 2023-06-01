package pg

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/foorester/cook/internal/infra/errors"
	"github.com/foorester/cook/internal/infra/sys"
)

type (
	DB struct {
		sys.Worker
		db *sqlx.DB
	}
)

const (
	name = "pg-db"
)

func NewDB(opts ...sys.Option) *DB {
	return &DB{
		Worker: sys.NewWorker(name, opts...),
		db:     nil,
	}
}

func (db *DB) Setup(ctx context.Context) error {
	return db.Connect()
}

func (db *DB) Connect() error {
	pgdb, err := sqlx.Open("postgres", db.connString())
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

func (db *DB) DB() *sqlx.DB {
	return db.db
}

func (db *DB) connString() (connString string) {
	cfg := db.Cfg()
	user := cfg.GetString("store.write.db.user")
	pass := cfg.GetString("store.write.db.pass")
	name := cfg.GetString("store.write.db.db")
	host := cfg.GetString("store.write.db.host")
	port := cfg.GetInt("store.write.db.port")
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d sslmode=require", user, pass, name, host, port)
}
