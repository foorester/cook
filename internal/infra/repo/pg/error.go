package pg

import "github.com/foorester/cook/internal/sys/errors"

var (
	NoConnectionError    = errors.NewError("no connection error")
	InvalidResourceIDErr = errors.NewError("invalid resource ID")
	UserNotFoundErr      = errors.NewError("user not found")
)
