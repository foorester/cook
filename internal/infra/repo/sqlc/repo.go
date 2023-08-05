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

func (cr *CookRepo) Conn(ctx context.Context) (*pgx.Conn, error) {
	return cr.db.PGXConn(ctx)
}

func (cr *CookRepo) GetUserByIDAndUsername(ctx context.Context, userID uuid.UUID, username string) (ok bool, user model.User, err error) {
	// WIP: Mock implementation
	uid, err := uuid.Parse("c4c109ad-f178-400a-b86d-3b0d548d852c")
	if err != nil {
		return false, user, InvalidResourceIDErr
	}

	ref := model.User{
		ID:       model.NewID(uid),
		Username: "johndoe",
	}

	if userID != ref.ID.Val() || username != ref.Username {
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

func (cr *CookRepo) GetBooks(ctx context.Context, userID uuid.UUID) (books []model.Book, err error) {
	ownerID, err := toPgUUID(userID)
	if err != nil {
		return books, errors.Wrap(err, "get books err")
	}

	conn, err := cr.Conn(ctx)
	if err != nil {
		return books, errors.Wrap(err, "create book err")
	}

	queries := New(conn)
	rows, err := queries.SelectAllBooks(ctx, ownerID)
	if err != nil {
		return books, errors.Wrap(err, "get books err")
	}

	for _, row := range rows {
		id, err := toID(row.ID)
		if err != nil {
			continue
		}

		book := model.Book{
			ID:          id,
			Name:        row.Name,
			Description: row.Description,
			Owner: model.User{
				// NOTE: eventually a variadic preload boolean parameter will be added
				// to the GetBooks method (`...preload bool`) to allow the caller
				// to preload the owner in the same query including all his values.
				ID: model.NewID(userID),
			},
			Audit: toAudit(row.CreatedAt, row.UpdatedAt),
		}

		books = append(books, book)
	}

	return books, nil
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

	queries := New(conn)
	_, err = queries.InsertBook(ctx, args)

	return err
}

func (cr *CookRepo) CreateRecipe(ctx context.Context, r model.Recipe) (err error) {
	return errors.New("not implemented yet")
}
