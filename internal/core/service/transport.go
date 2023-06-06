package service

import "github.com/foorester/cook/internal/core/model"

type (
	GetBooksReq struct {
	}

	GetBooksRes struct {
		Books []GetBookReq
	}
)

// GetBooksReq validation logic

type (
	SaveBookReq struct {
		ID   string
		Name string
	}

	SaveBookRes struct{}
)

// SaveBookReq validation logic

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
	SaveRecipeReq struct {
		ID   string
		Name string
	}

	SaveRecipeRes struct{}
)

func (req SaveRecipeReq) ToRecipe() model.Recipe {
	return model.Recipe{
		ID:   req.ID,
		Name: req.Name,
	}
}

// SaveRecipeReq validation logic

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
	SaveDirectionStepReq struct {
		Name        string
		Description string
		Duration    string
	}

	SaveDirectionStepRes struct{}
)

// SaveDirectionStepReq validation logic

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
	SaveIngredientReq struct {
		Name        string
		Description string
		Quantity    string
		Unit        string
	}

	SaveIngredientRes struct{}
)

// SaveIngredientReq validation logic

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
