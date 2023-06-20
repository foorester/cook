package service

import (
	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
)

type (
	UserReq struct {
		ID       uuid.UUID
		Username string
		Name     string
		Email    string
	}
)

type (
	GetBooksReq struct {
	}

	GetBooksRes struct {
		Books []GetBookReq
	}
)

// GetBooksReq validation logic

type (
	CreateBookReq struct {
		UserID      string
		Name        string
		Description string
	}
)

func (req CreateBookReq) ToBook() model.Book {
	return model.Book{
		Name:        req.Name,
		Description: req.Description,
	}
}

// CreateBookReq validation logic

type (
	DeleteBookReq struct {
		ID string
	}
)

// DeleteBookReq validation logic

type (
	GetBookReq struct {
		ID string
	}

	GetBookRes struct {
		ID   string
		Name string
	}
)

// GetBookReq validation logic

type (
	UpdateBookReq struct {
		ID   string
		Name string
	}
)

// UpdateBookReq validation logic

type (
	GetRecipesReq struct {
		ID string
	}

	GetRecipesRes struct {
		Recipes []GetRecipeRes
	}
)

// GetBookReq validation logic

type (
	CreateRecipeReq struct {
		UserID      string
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

// CreateRecipeReq validation logic

type (
	DeleteRecipeReq struct {
		ID string
	}
)

// DeleteReq validation logic

type (
	GetRecipeReq struct {
		ID string
	}

	GetRecipeRes struct {
		ID          string
		Name        string
		Ingredients []GetIngredientRes
		Steps       []GetDirectionStepRes
	}
)

// GetBookReq validation logic

type (
	UpdateRecipeReq struct {
		ID   string
		Name string
	}
)

// UpdateRecipeReq validation logic

type (
	GetDirectionStepsReq struct {
		ID string
	}

	GetDirectionStepsRes struct {
		Steps []GetDirectionStepRes
	}
)

// GetDirectionStepsReq validation logic

type (
	CreateDirectionStepReq struct {
		Name        string
		Description string
		Duration    string
	}
)

// CreateDirectionStepReq validation logic

type (
	GetDirectionStepReq struct {
		ID string
	}

	GetDirectionStepRes struct {
		ID          string
		Name        string
		Description string
		Duration    string
	}
)

// GetDirectionStepReq validation logic

type (
	UpdateDirectionStepReq struct {
		Name        string
		Description string
		Duration    string
	}
)

// UpdateDirectionStepReq validation logic

type (
	GetIngredientsReq struct {
		ID string
	}

	GetIngredientsRes struct {
		Steps []GetIngredientRes
	}
)

// GetIngredientsReq validation logic

type (
	CreateIngredientReq struct {
		Name        string
		Description string
		Quantity    string
		Unit        string
	}
)

// CreateIngredientReq validation logic

type (
	DeleteIngredientReq struct {
		ID string
	}
)

// DeleteIngredientReq validation logic

type (
	GetIngredientReq struct {
		ID string
	}

	GetIngredientRes struct {
		Id          string
		Name        string
		Description string
		Quantity    string
		Unit        string
	}
)

// GetIngredientReq validation logic

type (
	UpdateIngredientReq struct {
		Name        string
		Description string
		Quantity    string
		Unit        string
	}
)

// UpdateIngredientReq validation logic

// Misc
func toUUID(uuidStr string) (uid uuid.UUID, err error) {
	uid, err = uuid.Parse(uuidStr)
	if err != nil {
		return uid, InvalidID
	}

	return uid, nil
}
