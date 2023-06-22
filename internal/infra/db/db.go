package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/foorester/cook/internal/sys"
)

type (
	DB interface {
		sys.Core
		SQL
		PGX
		Mongo
		Connect(ctx context.Context) error
	}

	SQL interface {
		DBConn(ctx context.Context) (*sql.DB, error)
	}

	PGX interface {
		PGXConn(ctx context.Context) (*pgx.Conn, error)
	}

	Mongo interface {
		MongoConn(ctx context.Context) (*mongo.Client, error)
	}
)

type (
	UnimplementedSQL struct{}
)

func (u *UnimplementedSQL) DBConn(ctx context.Context) (*sql.DB, error) {
	return nil, errors.New("DBConn method is not implemented for this database")
}

type (
	UnimplementedPGX struct{}
)

func (u *UnimplementedPGX) PGXConn(ctx context.Context) (*pgx.Conn, error) {
	return nil, errors.New("PGXConn method is not implemented for this database")
}

type (
	UnimplementedNoSQL struct{}
)

func (u *UnimplementedNoSQL) MongoConn(ctx context.Context) (*mongo.Client, error) {
	return nil, errors.New("MongoConn method is not implemented for this database")
}
