package pg

import (
	"context"
	"fmt"

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
	pgConn, ok := pgxConn(db)
	if !ok {
		err = errors.Wrap("new cook repo", NoConnectionError)
		return nil, err
	}

	return &CookRepo{
		SimpleCore: sys.NewCore(name, opts...),
		db:         db,
		queries:    New(pgConn),
	}, nil
}

func (cr *CookRepo) Start(ctx context.Context) error {
	err := cr.db.Start(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s setup error", err)
		return errors.Wrap(msg, err)
	}

	return nil
}

func (cr *CookRepo) CreateBook(ctx context.Context, b model.Book) (err error) {
	args, err := toInsertBookParams(b)
	if err != nil {
		return errors.Wrap("create book err", err)
	}

	_, err = cr.queries.InsertBook(ctx, args)

	return err
}

func (cr *CookRepo) CreateRecipe(ctx context.Context, r model.Recipe) (err error) {
	return errors.NewError("not implemented yet")
}

func (cr *CookRepo) DB() (db any) {
	return cr.db.DB()
}

func pgxConn(db db.DB) (conn *pgx.Conn, ok bool) {
	conn, ok = db.DB().(*pgx.Conn)
	if !ok {
		return conn, false
	}

	return conn, true
}
