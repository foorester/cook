package transport

import (
	core "github.com/foorester/cook/internal/domain"
	"github.com/foorester/cook/internal/sys/config"
)

const (
	validationError = "Check fields with errors"
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
		exposeInternal: cfg.GetBool(config.Key.APIErrorExposeInt),
		valErrSet:      valErrSet,
		err:            err,
	}
}
