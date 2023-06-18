package mongo

import (
	"context"
	"database/sql"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	DB struct {
		sys.Core
		client *mongo.Client
		db     *mongo.Database
	}
)

const (
	name = "mongo-db"
)

func NewDB(opts ...sys.Option) *DB {
	return &DB{
		Core:   sys.NewCore(name, opts...),
		client: nil,
	}
}

func (db *DB) Start(ctx context.Context) error {
	return db.Connect()
}

func (db *DB) Connect() (err error) {
	connString := db.connString()
	db.client, err = mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return errors.Wrap("MongoDB connect client error", err)
	}

	ctx := context.TODO()

	err = db.client.Connect(ctx)
	if err != nil {
		return errors.Wrap("MongoDB connect error", err)
	}

	dbName := db.Cfg().GetString("store.write.db.db")
	db.db = db.client.Database(dbName)

	db.Log().Infof("%s database connected!", db.Name())
	return nil
}

func (db *DB) DB() any {
	return db.db
}

func (db *DB) mongoDB() (sqlDB *sql.DB, ok bool) {
	sqlDB, ok = db.DB().(*sql.DB)
	if !ok {
		return sqlDB, false
	}

	return sqlDB, true
}

func (db *DB) connString() (connString string) {
	cfg := db.Cfg()
	user := cfg.GetString("db.mongo.user")
	pass := cfg.GetString("db.mongo.pass")
	name := cfg.GetString("db.mongo.db")
	host := cfg.GetString("db.mongo.host")
	port := cfg.GetInt("db.mongo.port")
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, pass, host, port, name)
}
