package pg

import (
	"context"
	"database/sql"

	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	Tx struct {
		tx sql.Tx
	}

	TxContext struct {
		tx *sql.Tx
	}
)

func (db *DB) Begin(ctx context.Context) (db.TxContext, error) {
	pgTx, err := db.db.BeginTx(ctx, nil)

	return &TxContext{
		tx: pgTx,
	}, err

}

func (tc *TxContext) Commit(ctx context.Context) error {
	// WIP: Commit logic for postgres sql.Tx
	return errors.NotImplementedError
}

func (tc *TxContext) Rollback(ctx context.Context) error {
	// WIP: Rollback logic for postgres sql.Tx
	return errors.NotImplementedError
}
