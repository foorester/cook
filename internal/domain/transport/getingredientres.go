package transport

type (
	GetIngredientRes struct {
		ServiceRes
		Id          string
		Name        string
		Description string
		Quantity    string
		Unit        string
	}
)
