package sqlc

import "github.com/foorester/cook/internal/sys/errors"

var (
	NoConnectionError    = errors.New("no connection error")
	InvalidResourceIDErr = errors.New("invalid resource ID")
	UserNotFoundErr      = errors.New("user not found")
)
