package db

import (
	"github.com/jmoiron/sqlx"

	"github.com/foorester/cook/internal/sys"
)

type (
	DB interface {
		sys.Worker
		DB() *sqlx.DB
	}
)
