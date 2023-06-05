package port

import (
	"context"

	"github.com/foorester/cook/internal/core/model"
	"github.com/foorester/cook/internal/infra/db"
)

type (
	Repo interface {
		Tx(ctx context.Context) (tx db.Tx, err error)
	}

	RecipeRepo interface {
		Repo
		Save(ctx context.Context, recipe model.Recipe) error
	}
)
