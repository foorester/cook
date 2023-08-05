package transport

import (
	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
)

type (
	CreateBookReq struct {
		UserID      uuid.UUID
		Username    string
		Name        string
		Description string
	}
)

func (req CreateBookReq) ToBook() model.Book {
	return model.Book{
		Name:        req.Name,
		Description: req.Description,
	}
}
