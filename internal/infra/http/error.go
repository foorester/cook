package http

import "github.com/foorester/cook/internal/sys/errors"

type (
	HTTPError struct {
		errors.Err
	}
)

func NewError(msg string) HTTPError {
	return HTTPError{
		Err: errors.NewError(msg),
	}
}
