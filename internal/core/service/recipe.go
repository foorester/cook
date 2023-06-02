package service

import (
	"context"

	"github.com/foorester/cook/internal/core/model"
	"github.com/foorester/cook/internal/core/port"
)

type (
	RecipeService interface {
		SaveRecipe(ctx context.Context, r model.Recipe) error
	}

	Recipe struct {
		repo   port.RecipeRepo
		mailer port.Mailer
	}
)

func (rs *Recipe) SaveRecipe(ctx context.Context, r model.Recipe) error {
	panic("not implemented yet")
}
