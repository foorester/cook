package model

import "github.com/google/uuid"

type (
	Book struct {
		ID          uuid.UUID
		Name        string
		Description string
		Owner       User
	}
)
