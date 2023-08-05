package transport

type (
	GetRecipeRes struct {
		ServiceRes
		ID          string
		Name        string
		Ingredients []GetIngredientRes
		Steps       []GetDirectionStepRes
	}
)
