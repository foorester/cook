package sqlite

import (
	"database/sql"

	"github.com/foorester/cook/internal/infra/seed"
)

type (
	step struct {
		Index int64
		Name  string
		Seeds []seed.SeedFx
		tx    *sql.Tx
	}
)

func (s *step) Config(seed []seed.SeedFx) {
	s.Seeds = seed
}

func (s *step) GetIndex() (idx int64) {
	return s.Index
}

func (s *step) GetName() (name string) {
	return s.Name
}

func (s *step) GetSeeds() (seeds []seed.SeedFx) {
	return s.Seeds
}

func (s *step) SetTx(tx *sql.Tx) {
	s.tx = tx
}

func (s *step) GetTx() (tx *sql.Tx) {
	return s.tx
}
