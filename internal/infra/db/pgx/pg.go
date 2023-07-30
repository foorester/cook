package pgx

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/config"

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

var (
	cfgKey = config.Key
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
		return errors.Wrap(err, msg)
	}

	err = db.pool.Ping(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s ping connection error", db.Name())
		return errors.Wrap(err, msg)
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
	user := cfg.GetString(cfgKey.PgUser)
	pass := cfg.GetString(cfgKey.PgPass)
	name := cfg.GetString(cfgKey.PgDB)
	host := cfg.GetString(cfgKey.PgHost)
	port := cfg.GetInt(cfgKey.PgPort)
	schema := cfg.GetString(cfgKey.PgSchema)
	sslMode := cfg.GetBool(cfgKey.PgSSL)

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%d search_path=%s", user, pass, name, host, port, schema)
	db.Log().Debugf(connStr)

	if sslMode {
		connStr = connStr + " sslmode=enable"
	} else {
		connStr = connStr + " sslmode=disable"
	}

	return connStr
}
