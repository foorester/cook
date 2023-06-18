package model

import "github.com/google/uuid"

type (
	User struct {
		ID       uuid.UUID
		Username string
		Name     string
		Email    string
		Password string
	}
)
