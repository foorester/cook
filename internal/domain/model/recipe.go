package model

type (
	Recipe struct {
		ID          string
		Name        string
		Description string
		Book        Book
		Ingredients []Ingredient
		Steps       []Step
	}

	Ingredient struct {
		ID          string
		Name        string
		Description string
		Recipe      Recipe
		Qty         string // string for now, type may change
		Unit        string // string for now, type may change
	}

	Step struct {
		ID          string
		Name        string
		Description string
		Recipe      Recipe
		Duration    string // string for now, type may change
	}
)
