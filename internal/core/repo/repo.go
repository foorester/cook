package store

import (
	"context"

	"github.com/foorester/cook/internal/core/model"
)

type (
	Recipe interface {
		Save(ctx context.Context, recipe model.Recipe) error
	}
)
