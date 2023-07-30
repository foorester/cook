package seeding

import (
	"database/sql"

	"github.com/foorester/cook/internal/sys"
)

type (
	Seeder interface {
		sys.Core
		// Seed applies pending seeding
		Seed() error
		// SetAssetsPath sets the path form where the seeding are read
		SetAssetsPath(path string)
		// AssetsPath returns the path form where the seeding are read
		AssetsPath() string
	}
)

// SeedFx type alias
type SeedFx = func(tx *sql.Tx) error

type Exec interface {
	Config(seeds []SeedFx)
	GetIndex() (i int64)
	GetName() (name string)
	GetSeeds() (seedFxs []SeedFx)
	SetTx(tx *sql.Tx)
	GetTx() (tx *sql.Tx)
}
