package service

import (
	"context"

	"github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/domain/port"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	RecipeService interface {
		sys.Core
		Repo() port.CookRepo
		CreateBook(ctx context.Context, r model.Book) error
		CreateRecipe(ctx context.Context, r model.Recipe) error
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

func (rs *Recipe) CreateBook(ctx context.Context, req CreateBookReq) (errSet core.ValErrorSet, err error) {
	// Transport to Model
	book := req.ToBook()

	// Validate model
	v := NewBookValidator(book)

	err = v.ValidateForCreate()
	if err != nil {
		return v.Errors, err
	}

	// Persist it
	err = rs.Repo().CreateBook(ctx, book)
	if err != nil {
		return errSet, errors.Wrap("error creating recipe book", err)
	}

	// Send a message to bus
	return errSet, nil
}

func (rs *Recipe) CreateRecipe(ctx context.Context, req CreateRecipeReq) (errSet core.ValErrorSet, err error) {
	// Transport to Model
	recipe := req.ToRecipe()

	// Validate model
	v := NewRecipeValidator(recipe)

	err = v.ValidateForCreate()
	if err != nil {
		return v.Errors, err
	}

	// Persist it
	err = rs.Repo().CreateRecipe(ctx, recipe)
	if err != nil {
		return errSet, errors.Wrap("error creating recipe", err)
	}

	// Send a message to bus
	return errSet, nil
}

func (rs *Recipe) Repo() port.CookRepo {
	return rs.repo
}
