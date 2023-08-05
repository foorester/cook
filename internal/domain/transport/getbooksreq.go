package transport

import (
	"github.com/google/uuid"
)

type (
	GetBooksReq struct {
		UserID   uuid.UUID
		Username string
	}
)
