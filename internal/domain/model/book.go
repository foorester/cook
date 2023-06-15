package model

type (
	Book struct {
		ID          string
		Name        string
		Description string
		Owner       User
	}
)
