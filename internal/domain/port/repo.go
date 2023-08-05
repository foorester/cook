package port

import (
	"context"

	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/infra/db"
)

type (
	Repo interface {
		DB(ctx context.Context) (db db.DB)
	}

	CookRepo interface {
		Repo
		GetUser(ctx context.Context, userID string) (user model.User, err error)
		GetUserByIDAndUsername(ctx context.Context, userID uuid.UUID, username string) (ok bool, user model.User, err error)
		CreateBook(ctx context.Context, recipe model.Book) error
		GetBooks(ctx context.Context, userID uuid.UUID) (books []model.Book, err error)
		CreateRecipe(ctx context.Context, recipe model.Recipe) error
	}
)
