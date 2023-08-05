package transport

import (
	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
)

type (
	CreateRecipeReq struct {
		UserID      uuid.UUID
		Name        string
		Description string
		BookID      string
	}
)

func (req CreateRecipeReq) ToRecipe() model.Recipe {
	return model.Recipe{
		Name:        req.Name,
		Description: req.Description,
	}
}
