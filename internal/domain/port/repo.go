package port

import (
	"context"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/infra/db"
)

type (
	Repo interface {
		Tx(ctx context.Context) (tx db.Tx, err error)
	}

	CookRepo interface {
		Repo
		CreateBook(ctx context.Context, recipe model.Book) error
		CreateRecipe(ctx context.Context, recipe model.Recipe) error
		GetUser(ctx context.Context, userID string) (user model.User, err error)
	}
)
