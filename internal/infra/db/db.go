package db

import (
	"github.com/jmoiron/sqlx"

	"github.com/foorester/cook/internal/infra/sys"
)

type (
	DB interface {
		sys.Worker
		DB() *sqlx.DB
	}
)
