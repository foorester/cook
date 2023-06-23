package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys"

	"github.com/foorester/cook/internal/sys/errors"
)

type (
	DB struct {
		sys.Core
		pool *pgxpool.Pool
		db.UnimplementedSQL
		db.UnimplementedNoSQL
	}
)

const (
	name = "sqlc-db"
)

func NewDB(opts ...sys.Option) *DB {
	return &DB{
		Core: sys.NewCore(name, opts...),
	}
}

func (db *DB) Start(ctx context.Context) error {
	return db.Connect(ctx)
}

func (db *DB) Connect(ctx context.Context) (err error) {
	db.pool, err = pgxpool.New(ctx, db.connString())
	if err != nil {
		msg := fmt.Sprintf("%s connection error", db.Name())
		return errors.Wrap(msg, err)
	}

	err = db.pool.Ping(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", db.Name())
		return errors.Wrap(msg, err)
	}

	db.Log().Infof("%s database connected!", db.Name())
	return nil
}

func (db *DB) Conn(ctx context.Context) (conn *pgx.Conn, err error) {
	return db.PGXConn(ctx)
}

func (db *DB) PGXConn(ctx context.Context) (conn *pgx.Conn, err error) {
	poolConn, err := db.pool.Acquire(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire connection from pool: %w", err)
	}

	return poolConn.Conn(), nil
}

func (db *DB) connString() (connString string) {
	cfg := db.Cfg()
	user := cfg.GetString("db.pg.user")
	pass := cfg.GetString("db.pg.pass")
	name := cfg.GetString("db.pg.database")
	host := cfg.GetString("db.pg.host")
	port := cfg.GetInt("db.pg.port")
	schema := cfg.GetString("db.pg.schema")
	sslMode := cfg.GetBool("db.pg.sslmode")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d search_path=%s", user, pass, name, host, port, schema)
	db.Log().Infof(connStr)

	if sslMode {
		connStr = connStr + " sslmode=disable"
	} else {
		connStr = connStr + " sslmode=disable"
	}

	return connStr
}
