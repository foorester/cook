package pg

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/infra/db"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	CookRepo struct {
		*sys.SimpleCore
		db      db.DB
		queries *Queries
	}
)

const (
	name = "cook-repo"
)

func NewCookRepo(db db.DB, opts ...sys.Option) (cr *CookRepo, err error) {
	return &CookRepo{
		SimpleCore: sys.NewCore(name, opts...),
		db:         db,
	}, nil
}

func (cr *CookRepo) Setup(ctx context.Context) (err error) {
	err = cr.db.Connect(ctx)
	if err != nil {
		err = errors.Wrap("cook repo setup error", err)
		return err
	}

	return nil
}

func (cr *CookRepo) DB(ctx context.Context) (db db.DB) {
	return cr.db
}

func (cr *CookRepo) CreateBook(ctx context.Context, b model.Book) (err error) {
	b.GenID()

	args, err := toInsertBookParams(b)
	if err != nil {
		return errors.Wrap("create book err", err)
	}

	conn, err := cr.Conn(ctx)
	if err != nil {
		return errors.Wrap("create book err", err)
	}

	queries := NewQueries(conn)
	_, err = queries.InsertBook(ctx, args)

	return err
}

func (cr *CookRepo) CreateRecipe(ctx context.Context, r model.Recipe) (err error) {
	return errors.NewError("not implemented yet")
}

func (cr *CookRepo) Conn(ctx context.Context) (*pgx.Conn, error) {
	return cr.db.PGXConn(ctx)
}

func (cr *CookRepo) GetUser(ctx context.Context, userID string) (user model.User, err error) {
	// WIP: Mock implementation
	ref := "c4c109ad-f178-400a-b86d-3b0d548d852c"

	uid, err := uuid.Parse(userID)
	if err != nil {
		return user, InvalidResourceIDErr
	}

	if userID == ref {
		return model.User{
			ID:       uid,
			Username: "johndoe",
			Name:     "John Doe",
			Email:    "john.doe@localhost.com",
		}, nil

	}

	return user, UserNotFoundErr
}
