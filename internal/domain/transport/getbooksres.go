package transport

import (
	core "github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/sys/config"
)

type (
	GetBooksRes struct {
		ServiceRes
		Books []GetBookRes
	}
)

func NewGetBooksRes(valErrSet core.ValErrorSet, err error, cfg *config.Config) GetBooksRes {
	return GetBooksRes{
		ServiceRes: NewServiceRes(valErrSet, err, cfg),
	}
}

func (gbr *GetBooksRes) SetBooks(books []model.Book) {
	gbrs := make([]GetBookRes, 0, len(books))
	for _, book := range books {
		gbrs = append(gbrs, GetBookRes{
			ID:          book.ID.String(),
			Name:        book.Name,
			Description: book.Description,
			OwnerID:     book.Owner.ID.String(),
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
		})
	}

	gbr.Books = gbrs
}
