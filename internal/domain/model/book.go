package model

type (
	Book struct {
		ID
		Name        string
		Description string
		Owner       User
		Audit
	}
)
