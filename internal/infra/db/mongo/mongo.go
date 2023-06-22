package mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	DB struct {
		sys.Core
		client *mongo.Client
		mongo  *mongo.Database
		db.UnimplementedSQL
		db.UnimplementedPGX
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
	return db.Connect(ctx)
}

func (db *DB) Connect(ctx context.Context) (err error) {
	connString := db.connString()
	db.client, err = mongo.NewClient(options.Client().ApplyURI(connString))
	if err != nil {
		return errors.Wrap("MongoDB connect client error", err)
	}

	err = db.client.Connect(ctx)
	if err != nil {
		return errors.Wrap("MongoDB connect error", err)
	}

	dbName := db.Cfg().GetString("store.write.db.mongo")
	db.mongo = db.client.Database(dbName)

	db.Log().Infof("%s database connected!", db.Name())
	return nil
}

func (db *DB) MongoConn(ctx context.Context) (*mongo.Client, error) {
	return db.mongo.Client(), nil
}

func (db *DB) connString() (connString string) {
	cfg := db.Cfg()
	user := cfg.GetString("db.mongo.user")
	pass := cfg.GetString("db.mongo.pass")
	name := cfg.GetString("db.mongo.database")
	host := cfg.GetString("db.mongo.host")
	port := cfg.GetInt("db.mongo.port")
	return fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", user, pass, host, port, name)
}
