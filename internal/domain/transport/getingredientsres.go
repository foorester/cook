package transport

type (
	GetIngredientsRes struct {
		ServiceRes
		Steps []GetIngredientRes
	}
)
