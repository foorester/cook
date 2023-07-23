package migration

import (
	"database/sql"

	"github.com/foorester/cook/internal/infra/migration/sqlite"
	"github.com/foorester/cook/internal/sys"
)

type (
	Migrator interface {
		sys.Core
		// Migrate applies pending seeding
		Migrate() error
		// Rollback reverts from one to N seeding already applied
		Rollback(steps ...int) error
		// RollbackAll reverts all seeding allready applied
		RollbackAll() error
		// SoftReset apply all seeding again after rolling back all seeding.
		SoftReset() error
		// Reset apply all seeding again after dropping the database and recreating it
		Reset() error
		// SetAssetsPath sets the path form where the seeding are read
		SetAssetsPath(path string)
		// AssetsPath returns the path form where the seeding are read
		AssetsPath() string
	}
)

type Exec interface {
	Config(up sqlite.MigFx, down sqlite.MigFx)
	GetIndex() (i int64)
	GetName() (name string)
	GetUp() (up sqlite.MigFx)
	GetDown() (down sqlite.MigFx)
	SetTx(tx *sql.Tx)
	GetTx() (tx *sql.Tx)
}
