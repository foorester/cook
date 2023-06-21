package service

import (
	"context"

	"github.com/foorester/cook/internal/domain/port"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	RecipeService interface {
		sys.Core
		Repo() port.CookRepo
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

	// Validate model
	v := NewBookValidator(book)

	err := v.ValidateForCreate()
	if err != nil {
		return NewCreateBookRes(v.Errors, err, rs.Cfg())
	}

	// Set Owner
	user, err := rs.Repo().GetUser(ctx, req.UserID)
	if err != nil {
		err = errors.Wrap("create book error", err)
		return NewCreateBookRes(nil, err, rs.Cfg())
	}

	book.Owner = user

	// Persist it
	err = rs.Repo().CreateBook(ctx, book)
	if err != nil {
		err = errors.Wrap("create book error", err)
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
		err = errors.Wrap("create recipe error", err)
		return NewCreateRecipeRes(nil, err, rs.Cfg())
	}

	// Send a message to bus

	return NewCreateRecipeRes(nil, nil, nil)
}

func (rs *Recipe) Repo() port.CookRepo {
	return rs.repo
}
