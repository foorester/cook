package transport

import (
	core "github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/sys/config"
)

type (
	CreateBookRes struct {
		ServiceRes
		ID string
	}
)

func NewCreateBookRes(valErrSet core.ValErrorSet, err error, cfg *config.Config) CreateBookRes {
	return CreateBookRes{
		ServiceRes: NewServiceRes(valErrSet, err, cfg),
	}
}
