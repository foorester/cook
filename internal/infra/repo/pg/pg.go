package pg

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	RecipeRepo struct {
		*sys.SimpleCore
		db db.DB
	}
)

const (
	name = "write-store"
)

func NewRecipeRepo(db db.DB, opts ...sys.Option) *RecipeRepo {
	return &RecipeRepo{
		SimpleCore: sys.NewCore(name, opts...),
		db:         db,
	}
}

func (rr *RecipeRepo) Start(ctx context.Context) error {
	err := rr.db.Start(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s setup error", err)
		return errors.Wrap(msg, err)
	}

	return nil
}

func (rr *RecipeRepo) DB() (db any) {
	return rr.db.DB()
}

func (rr *RecipeRepo) PgDB() (db *sqlx.DB, ok bool) {
	db, ok = rr.DB().(*sqlx.DB)
	if !ok {
		return db, false
	}

	return db, true
}

func (rr *RecipeRepo) CreateBook(ctx context.Context, r model.Book) (err error) {
	books := []model.Book{r}

	db, ok := rr.PgDB()
	if !ok {
		return NoConnectionError
	}

	// TODO: Update SQL Statement, there is a better way to do this insertion
	_, err = db.NamedExec(`INSERT INTO books (id, name, description) VALUES (:id, :name, :description)`, books)

	if err != nil {
		return errors.Wrap("save recipe book error", err)
	}

	return nil
}

func (rr *RecipeRepo) CreateRecipe(ctx context.Context, r model.Recipe) (err error) {
	recipes := []model.Recipe{r}

	db, ok := rr.PgDB()
	if !ok {
		return NoConnectionError
	}

	// TODO: Update SQL Statement, there is a better way to do this insertion
	_, err = db.NamedExec(`INSERT INTO recipes (id, name) VALUES (:id, :name)`, recipes)

	if err != nil {
		return errors.Wrap("save recipe error", err)
	}

	return nil
}
