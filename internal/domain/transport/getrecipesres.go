package transport

type (
	GetRecipesRes struct {
		ServiceRes
		Recipes []GetRecipeRes
	}
)
