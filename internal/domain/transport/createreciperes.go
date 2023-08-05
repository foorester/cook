package transport

import (
	core "github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/sys/config"
)

type (
	CreateRecipeRes struct {
		ServiceRes
		ID string
	}
)

func NewCreateRecipeRes(valErrSet core.ValErrorSet, err error, cfg *config.Config) CreateRecipeRes {
	return CreateRecipeRes{
		ServiceRes: NewServiceRes(valErrSet, err, cfg),
	}
}
