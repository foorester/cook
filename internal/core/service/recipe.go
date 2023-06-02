package service

import (
	"context"

	"github.com/foorester/cook/internal/core/model"
	"github.com/foorester/cook/internal/core/port"
	"github.com/foorester/cook/internal/sys"
)

type (
	RecipeService interface {
		sys.Worker
		SaveRecipe(ctx context.Context, r model.Recipe) error
	}

	Recipe struct {
		sys.Worker
		repo   port.RecipeRepo
		mailer port.Mailer
	}
)

const (
	name = "repo-service"
)

func NewService(rr port.RecipeRepo, opts ...sys.Option) *Recipe {
	return &Recipe{
		Worker: sys.NewWorker(name, opts...),
		repo:   rr,
		mailer: nil, // Interface not implemented yet
	}
}

func (rs *Recipe) SaveRecipe(ctx context.Context, r model.Recipe) error {
	panic("not implemented yet")
}
