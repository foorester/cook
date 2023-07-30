package seeding

import (
	"database/sql"
)

type (
	Step struct {
		Index int64
		Name  string
		Seeds []SeedFx
		tx    *sql.Tx
	}
)

func (s *Step) Config(seed []SeedFx) {
	s.Seeds = seed
}

func (s *Step) GetIndex() (idx int64) {
	return s.Index
}

func (s *Step) GetName() (name string) {
	return s.Name
}

func (s *Step) GetSeeds() (seeds []SeedFx) {
	return s.Seeds
}

func (s *Step) SetTx(tx *sql.Tx) {
	s.tx = tx
}

func (s *Step) GetTx() (tx *sql.Tx) {
	return s.tx
}
