package transport

import "github.com/google/uuid"

type (
	UserReq struct {
		ID       uuid.UUID
		Username string
		Name     string
		Email    string
	}
)
