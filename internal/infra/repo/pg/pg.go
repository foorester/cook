package pg

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/foorester/cook/internal/core/model"
	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/infra/errors"
	"github.com/foorester/cook/internal/infra/sys"
)

type (
	RecipeRepo struct {
		sys.Worker
		db *db.DB
	}
)

const (
	name = "write-store"
)

func NewRecipeRepo(db *db.DB, opts ...sys.Option) *RecipeRepo {
	return &RecipeRepo{
		Worker: sys.NewWorker(name, opts...),
		db:     db,
	}
}

func (rr *RecipeRepo) Setup(ctx context.Context) error {
	err := rr.db.Setup(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s setup error", err)
		return errors.Wrap(msg, err)
	}

	return nil
}

func (rr *RecipeRepo) Save(ctx context.Context, r model.Recipe) (err error) {
	recipes := []model.Recipe{r}

	_, err = rr.db.NamedExec(`INSERT INTO recipes (id, name) VALUES (:id, :name)`, recipes)

	if err != nil {
		return errors.Wrap("save recipe error", err)
	}

	return nil
}
