package sqlc

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
		err = errors.Wrap(err, "cook repo setup error")
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
		return errors.Wrap(err, "create book err")
	}

	conn, err := cr.Conn(ctx)
	if err != nil {
		return errors.Wrap(err, "create book err")
	}

	queries := NewQueries(conn)
	_, err = queries.InsertBook(ctx, args)

	return err
}

func (cr *CookRepo) CreateRecipe(ctx context.Context, r model.Recipe) (err error) {
	return errors.New("not implemented yet")
}

func (cr *CookRepo) Conn(ctx context.Context) (*pgx.Conn, error) {
	return cr.db.PGXConn(ctx)
}

func (cr *CookRepo) IsValidUser(ctx context.Context, userID, username string) (ok bool, user model.User, err error) {
	// WIP: Mock implementation
	uid, err := uuid.Parse("c4c109ad-f178-400a-b86d-3b0d548d852c")
	if err != nil {
		return false, user, InvalidResourceIDErr
	}

	ref := model.User{
		ID:       model.NewID(uid),
		Username: "johndoe",
	}

	_, err = uuid.Parse(userID)
	if err != nil {
		return false, user, InvalidResourceIDErr
	}

	if userID != ref.ID.String() || username != ref.Username {
		return false, user, nil
	}

	return true, ref, nil
}

func (cr *CookRepo) GetUser(ctx context.Context, userID string) (user model.User, err error) {
	// WIP: Mock implementation
	uid, err := uuid.Parse("c4c109ad-f178-400a-b86d-3b0d548d852c")
	if err != nil {
		return user, InvalidResourceIDErr
	}

	uid, err = uuid.Parse(userID)
	if err != nil {
		return user, InvalidResourceIDErr
	}

	if userID == uid.String() {
		return model.User{
			ID:       model.NewID(uid),
			Username: "johndoe",
			Name:     "John Doe",
			Email:    "john.doe@localhost.com",
		}, nil

	}

	return user, UserNotFoundErr
}
