package model

type (
	Recipe struct {
		ID
		Name        string
		Description string
		Book        Book
		Ingredients []Ingredient
		Steps       []Step
		Audit
	}

	Ingredient struct {
		ID
		Name        string
		Description string
		Recipe      Recipe
		Qty         string // string for now, type may change
		Unit        string // string for now, type may change
		Audit
	}

	Step struct {
		ID
		Name        string
		Description string
		Recipe      Recipe
		Duration    string // string for now, type may change
		Audit
	}
)
