package service

import (
	core "github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/sys/config"
)

const (
	validationError = "Check fields with errors"
)

const (
	exposeIntKey = "api.errors.expose.internal"
)

type (
	ServiceRes struct {
		msg            string           // Human readable message exposed to client
		exposeInternal bool             // Expose internal error to te client flag
		valErrSet      core.ValErrorSet // Properties validation errors
		err            error            // Internal error
	}
)

func (sr *ServiceRes) Msg() string {
	return sr.msg
}

func (sr *ServiceRes) ValidationErrors() core.ValErrorSet {
	return sr.valErrSet
}

func (sr *ServiceRes) Err() error {
	if sr.exposeInternal {
		return sr.err
	}

	return nil
}

func NewServiceRes(valErrSet core.ValErrorSet, err error, cfg *config.Config) ServiceRes {
	return ServiceRes{
		msg:            validationError,
		exposeInternal: cfg.GetBool(exposeIntKey),
		valErrSet:      valErrSet,
		err:            err,
	}
}

type (
	CreateBookRes struct {
		ServiceRes
	}
)

func NewCreateBookRes(valErrSet core.ValErrorSet, err error, cfg *config.Config) CreateBookRes {
	return CreateBookRes{
		ServiceRes: NewServiceRes(valErrSet, err, cfg),
	}
}

type (
	CreateRecipeRes struct {
		ServiceRes
	}
)

func NewCreateRecipeRes(valErrSet core.ValErrorSet, err error, cfg *config.Config) CreateRecipeRes {
	return CreateRecipeRes{
		ServiceRes: NewServiceRes(valErrSet, err, cfg),
	}
}
