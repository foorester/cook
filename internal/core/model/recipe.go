package model

type (
	Recipe struct {
		ID          string
		Name        string
		Description string
		Ingredients []Ingredient
		Steps       []Step
	}

	Ingredient struct {
		Id   string
		Name string
		Qty  string // string for now, type may change
		Unit string // string for now, type may change
	}

	Step struct {
		Id          string
		Name        string
		Description string
		Duration    string // string for now, type may change
	}
)
