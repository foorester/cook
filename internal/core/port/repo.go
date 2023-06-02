package port

import (
	"context"

	"github.com/foorester/cook/internal/core/model"
)

type (
	RecipeRepo interface {
		Save(ctx context.Context, recipe model.Recipe) error
	}
)
