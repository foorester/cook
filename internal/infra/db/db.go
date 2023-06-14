package db

import (
	"context"

	"github.com/foorester/cook/internal/sys"
)

type (
	DB interface {
		sys.Core
		// DB - NOTE: See if there is a common set of features that can define what is returned by
		// this function. For the moment it will return any (interface{}).
		// Not ideal, forces type assertion when using specific implementation.
		DB() any
		Tx
	}

	Tx interface {
		Begin(ctx context.Context) (TxContext, error)
	}

	TxContext interface {
		Commit(ctx context.Context) error
		Rollback(ctx context.Context) error
	}
)
