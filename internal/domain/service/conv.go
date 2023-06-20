package service

import (
	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
)

type (
	User struct {
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
		Name        string
		Description string
	}

	CreateBookRes struct{}
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

	DeleteBookRes struct{}
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

	UpdateBookRes struct {
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
		Name string
	}

	CreateRecipeRes struct{}
)

func (req CreateRecipeReq) ToRecipe() model.Recipe {
	return model.Recipe{
		Name: req.Name,
	}
}

// CreateRecipeReq validation logic

type (
	DeleteRecipeReq struct {
		ID string
	}

	DeleteRecipeRes struct{}
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

	UpdateRecipeRes struct {
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

	CreateDirectionStepRes struct{}
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

	UpdateDirectionStepRes struct {
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

	CreateIngredientRes struct{}
)

// CreateIngredientReq validation logic

type (
	DeleteIngredientReq struct {
		ID string
	}

	DeleteIngredientRes struct{}
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

	UpdateIngredientRes struct {
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
