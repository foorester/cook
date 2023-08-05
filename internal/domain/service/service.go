package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/domain/port"
	t "github.com/foorester/cook/internal/domain/transport"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	RecipeService interface {
		sys.Core
		GetBooks(ctx context.Context, m t.GetBooksReq) t.GetBooksRes
		CreateBook(ctx context.Context, m t.CreateBookReq) t.CreateBookRes
		CreateRecipe(ctx context.Context, m t.CreateRecipeReq) t.CreateRecipeRes
	}

	Recipe struct {
		*sys.SimpleCore
		repo   port.CookRepo
		mailer port.Mailer
	}
)

const (
	name = "repo-service"
)

func NewService(rr port.CookRepo, opts ...sys.Option) *Recipe {
	return &Recipe{
		SimpleCore: sys.NewCore(name, opts...),
		repo:       rr,
		mailer:     nil, // Interface not implemented yet
	}
}

func (rs *Recipe) GetBooks(ctx context.Context, req t.GetBooksReq) (res t.GetBooksRes) {
	_, err := rs.validateUser(ctx, req.UserID, req.Username)
	if err != nil {
		err = errors.Wrap(err, "get books error")
		return t.NewGetBooksRes(nil, err, rs.Cfg())
	}

	books, err := rs.Repo().GetBooks(ctx, req.UserID)
	if err != nil {
		err = errors.Wrap(err, "create req error")
		return t.NewGetBooksRes(nil, err, rs.Cfg())
	}

	res.SetBooks(books)

	return t.NewGetBooksRes(nil, nil, nil)
}

func (rs *Recipe) CreateBook(ctx context.Context, req t.CreateBookReq) (res t.CreateBookRes) {
	book := req.ToBook()

	user, err := rs.validateUser(ctx, req.UserID, req.Username)
	if err != nil {
		err = errors.Wrap(err, "create book error")
		return t.NewCreateBookRes(nil, err, rs.Cfg())
	}

	v := NewBookValidator(book)
	err = v.ValidateForCreate()
	if err != nil {
		return t.NewCreateBookRes(v.Errors, err, rs.Cfg())
	}

	book.Owner = user

	err = rs.Repo().CreateBook(ctx, book)
	if err != nil {
		err = errors.Wrap(err, "create book error")
		return t.NewCreateBookRes(nil, err, rs.Cfg())
	}

	return t.NewCreateBookRes(nil, nil, nil)
}

func (rs *Recipe) CreateRecipe(ctx context.Context, req t.CreateRecipeReq) (res t.CreateRecipeRes) {
	recipe := req.ToRecipe()

	v := NewRecipeValidator(recipe)

	err := v.ValidateForCreate()
	if err != nil {
		return t.NewCreateRecipeRes(v.Errors, err, rs.Cfg())
	}

	err = rs.Repo().CreateRecipe(ctx, recipe)
	if err != nil {
		err = errors.Wrap(err, "create recipe error")
		return t.NewCreateRecipeRes(nil, err, rs.Cfg())
	}

	// Send a message to bus

	return t.NewCreateRecipeRes(nil, nil, nil)
}

func (rs *Recipe) Repo() port.CookRepo {
	return rs.repo
}

func (rs *Recipe) Start(ctx context.Context) error {
	db := rs.repo.DB(ctx)

	err := db.Start(ctx)
	if err != nil {
		msg := fmt.Sprintf("%s start error", rs.Name())
		return errors.Wrap(err, msg)
	}

	return nil
}

func (rs *Recipe) validateUser(ctx context.Context, userID uuid.UUID, username string) (user model.User, err error) {
	ok, user, err := rs.Repo().GetUserByIDAndUsername(ctx, userID, username)
	if err != nil {
		return user, err
	}

	if !ok {
		return user, errors.New("invalid username")
	}

	return user, nil
}
