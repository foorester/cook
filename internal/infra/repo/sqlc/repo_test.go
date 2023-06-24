package sqlc_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"

	"github.com/foorester/cook/internal/domain/model"
	"github.com/foorester/cook/internal/infra/db/pgx"
	"github.com/foorester/cook/internal/infra/repo/sqlc"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/log"
)

var cfgKey = config.Key

func TestCreateBook(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	c := testCfg(t)
	l := testLog(t)

	opts := []sys.Option{
		sys.WithConfig(c),
		sys.WithLogger(l),
	}

	dbase := pgx.NewDB(opts...)
	cookRepo, err := sqlc.NewCookRepo(dbase, opts...)
	ctx := context.Background()

	book := model.Book{
		ID:          model.NewID(uuid.New()),
		Name:        "Sample Book",
		Description: "This is a sample book",
		Owner: model.User{
			ID:       uuid.New(),
			Username: "johndoe",
			Name:     "John Doe",
			Email:    "john.doe@example.com",
			Password: "password123",
		},
	}

	err = cookRepo.Setup(ctx)
	if err != nil {
		t.Errorf("Create book test setup error: %s", err.Error())
	}

	err = cookRepo.CreateBook(ctx, book)
	if err != nil {
		t.Errorf("Create book test error: %s", err.Error())
	}

	// Add more real use-case assertions
	t.Fatal("not implemented yet")
}

func testCfg(t *testing.T) *config.Config {
	cfg := config.Config{}
	cfg.SetValues(map[string]string{
		cfgKey.APIServerHost:     "localhost",
		cfgKey.APIServerPort:     "8080",
		cfgKey.APIServerTimeout:  "10",
		cfgKey.APIErrorExposeInt: "true",

		// Postgres

		cfgKey.PgUser:   "cook",
		cfgKey.PgPass:   "cook",
		cfgKey.PgDB:     "foorester",
		cfgKey.PgHost:   "pg",
		cfgKey.PgPort:   "5432",
		cfgKey.PgSchema: "cook",
		cfgKey.PgSSL:    "false",

		// Mongo

		cfgKey.MongoUser: "cook",
		cfgKey.MongoPass: "cook",
		cfgKey.MongoDB:   "cook",
		cfgKey.MongoHost: "mongo",
		cfgKey.MongoPort: "27020",
	})

	return &cfg
}

func testLog(t *testing.T) *log.TestLogger {
	return log.NewTestLogger("debug")
}
