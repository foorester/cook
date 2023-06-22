package pg

import (
	"errors"
)

type (
	ctxKey string
)

const (
	TxKey = ctxKey("tx")
)

var (
	TxNotFoundError = errors.New("no transaction found in context")
)
