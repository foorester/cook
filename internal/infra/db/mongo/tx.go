package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	Tx struct {
		session mongo.Session
	}

	TxContext struct {
		session mongo.Session
	}
)

func (db *DB) BeginTx(ctx context.Context) (db.TxContext, error) {
	session, err := db.client.StartSession()
	if err != nil {
		return nil, err
	}
	err = session.StartTransaction()
	if err != nil {
		session.EndSession(context.Background())
		return nil, err
	}
	return &TxContext{session: session}, nil
}

func (tc *TxContext) Commit(ctx context.Context) error {
	// WIP: Commit logic for mongo
	return errors.NotImplementedError
}

func (tc *TxContext) Rollback(ctx context.Context) error {
	// WIP: Rollback logic for mongo
	return errors.NotImplementedError
}
