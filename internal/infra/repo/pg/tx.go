package pg

import (
	"context"
	"errors"

	"github.com/foorester/cook/internal/infra/db"
)

type (
	ctxKey string
)

const (
	TxKey = ctxKey("tx")
)

var (
	TxNotFoundError = errors.New("no transaction found in context")
)

func (cr *CookRepo) Tx(ctx context.Context) (db.Tx, error) {
	tx, ok := ctx.Value(TxKey).(db.Tx)
	if !ok {
		return nil, TxNotFoundError
	}

	return tx, nil
}
