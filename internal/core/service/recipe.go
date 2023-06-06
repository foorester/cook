package service

import (
	"context"

	"github.com/foorester/cook/internal/core"
	"github.com/foorester/cook/internal/core/model"
	"github.com/foorester/cook/internal/core/port"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/errors"
)

type (
	RecipeService interface {
		sys.Worker
		Repo() port.RecipeRepo
		SaveRecipe(ctx context.Context, r model.Recipe) error
	}

	Recipe struct {
		*sys.BaseWorker
		repo   port.RecipeRepo
		mailer port.Mailer
	}
)

const (
	name = "repo-service"
)

func NewService(rr port.RecipeRepo, opts ...sys.Option) *Recipe {
	return &Recipe{
		BaseWorker: sys.NewWorker(name, opts...),
		repo:       rr,
		mailer:     nil, // Interface not implemented yet
	}
}

func (rs *Recipe) SaveRecipe(ctx context.Context, req SaveRecipeReq) (errSet core.ValErrorSet, err error) {
	// Transport to Model
	recipe := req.ToRecipe()

	// Validate model
	v := NewRecipeValidator(recipe)

	err = v.ValidateForSave()
	if err != nil {
		return v.Errors, err
	}

	// Persist it
	err = rs.Repo().Save(ctx, recipe)
	if err != nil {
		return errSet, errors.Wrap("error saving recipe", err)
	}

	// Send a message to bus
	return errSet, nil
}

func (rs *Recipe) Repo() port.RecipeRepo {
	return rs.repo
}
