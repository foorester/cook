package sqlite

import (
	"database/sql"

	"github.com/foorester/cook/internal/infra/migration"
)

type (
	step struct {
		Index int64
		Name  string
		Up    migration.MigFx
		Down  migration.MigFx
		tx    *sql.Tx
	}
)

func (s *step) Config(up migration.MigFx, down migration.MigFx) {
	s.Up = up
	s.Down = down
}

func (s *step) GetIndex() (idx int64) {
	return s.Index
}

func (s *step) GetName() (name string) {
	return s.Name
}

func (s *step) GetUp() (up migration.MigFx) {
	return s.Up
}

func (s *step) GetDown() (down migration.MigFx) {
	return s.Down
}

func (s *step) SetTx(tx *sql.Tx) {
	s.tx = tx
}

func (s *step) GetTx() (tx *sql.Tx) {
	return s.tx
}
