package service

import (
	"context"
	"fmt"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/domain/port"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	RecipeService interface {
		sys.Core
		CreateBook(ctx context.Context, m CreateBookReq) CreateBookRes
		CreateRecipe(ctx context.Context, m CreateRecipeReq) CreateRecipeRes
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

func (rs *Recipe) CreateBook(ctx context.Context, req CreateBookReq) (res CreateBookRes) {
	// Transport to Model
	book := req.ToBook()

	// Owner validation
	user, err := rs.validateUser(ctx, req.UserID, req.Username)
	if err != nil {
		err = errors.Wrap(err, "create book error")
		return NewCreateBookRes(nil, err, rs.Cfg())
	}

	// Model validation
	v := NewBookValidator(book)

	err = v.ValidateForCreate()
	if err != nil {
		return NewCreateBookRes(v.Errors, err, rs.Cfg())
	}

	// Set Owner
	book.Owner = user

	// Persist it
	err = rs.Repo().CreateBook(ctx, book)
	if err != nil {
		err = errors.Wrap(err, "create book error")
		return NewCreateBookRes(nil, err, rs.Cfg())
	}

	return NewCreateBookRes(nil, nil, nil)
}

func (rs *Recipe) CreateRecipe(ctx context.Context, req CreateRecipeReq) (res CreateRecipeRes) {
	// Transport to Model
	recipe := req.ToRecipe()

	// Validate model
	v := NewRecipeValidator(recipe)

	err := v.ValidateForCreate()
	if err != nil {
		return NewCreateRecipeRes(v.Errors, err, rs.Cfg())
	}

	// Persist it
	err = rs.Repo().CreateRecipe(ctx, recipe)
	if err != nil {
		err = errors.Wrap(err, "create recipe error")
		return NewCreateRecipeRes(nil, err, rs.Cfg())
	}

	// Send a message to bus

	return NewCreateRecipeRes(nil, nil, nil)
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

func (rs *Recipe) validateUser(ctx context.Context, userID, username string) (user model.User, err error) {
	ok, user, err := rs.Repo().IsValidUser(ctx, userID, username)
	if err != nil {
		return user, err
	}

	if !ok {
		return user, errors.New("invalid username")
	}

	return user, nil
}
