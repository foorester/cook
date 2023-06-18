package pg

import "github.com/foorester/cook/internal/sys/errors"

var (
	NoConnectionError = errors.NewError("no connection error")
	InvalidID         = errors.NewError("invalid ID")
)
