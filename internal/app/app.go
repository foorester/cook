package app

import (
	"context"
	"embed"
	"fmt"
	"sync"

	"github.com/foorester/cook/internal/domain/service"
	"github.com/foorester/cook/internal/infra/db/pgx"
	http2 "github.com/foorester/cook/internal/infra/http"
	"github.com/foorester/cook/internal/infra/migration"
	mig "github.com/foorester/cook/internal/infra/migration/pg"
	pgxr "github.com/foorester/cook/internal/infra/repo/sqlc"
	"github.com/foorester/cook/internal/infra/seeding"
	seed "github.com/foorester/cook/internal/infra/seeding/pg"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/errors"
	"github.com/foorester/cook/internal/sys/log"
)

type App struct {
	sync.Mutex
	sys.Core
	opts       []sys.Option
	migFs      embed.FS
	seedFs     embed.FS
	supervisor sys.Supervisor
	http       *http2.Server
	svc        service.RecipeService
	migrator   migration.Migrator
	seeder     seeding.Seeder
}

func NewApp(name string, log log.Logger) (app *App, err error) {
	cfg, err := config.NewConfig(name).Load()
	if err != nil {
		return nil, errors.Wrap(err, launchError)
	}

	opts := []sys.Option{
		sys.WithConfig(cfg),
		sys.WithLogger(log),
	}

	app = &App{
		Core: sys.NewCore(name, opts...),
		opts: opts,
	}

	return app, nil
}

func (app *App) SetMigrationsFs(fs embed.FS) {
	app.migFs = fs
}

func (app *App) SetSeedingFs(fs embed.FS) {
	app.seedFs = fs
}

func (app *App) Run() (err error) {
	ctx := context.Background()

	err = app.Setup(ctx)
	if err != nil {
		return errors.Wrap(err, runError)
	}

	return app.Start(ctx)
}

func (app *App) Setup(ctx context.Context) error {
	app.EnableSupervisor()

	// Migrator
	app.migrator = mig.NewMigrator(app.migFs, app.opts...)

	// Seeder
	app.seeder = seed.NewSeeder(app.seedFs, app.opts...)

	// Databases
	dbase := pgx.NewDB(app.opts...)

	// Repos
	repo, err := pgxr.NewCookRepo(dbase, app.opts...)
	if err != nil {
		return err
	}

	// Services
	app.svc = service.NewService(repo, app.opts...)

	// HTTP Server
	app.http = http2.NewServer(app.svc, app.opts...)
	app.SetupRoutes(ctx)
	app.SetupProbes(ctx)

	// gRPC servers

	// Event bus

	// WIP: to avoid unused var message
	//app.Log().Debugf("Repo: %v", repo)
	//app.Log().Debugf("Service: %v", app.svc)

	return nil
}

func (app *App) Start(ctx context.Context) error {
	app.Log().Infof("%s starting...", app.Name())
	defer app.Log().Infof("%s stopped", app.Name())

	var err error

	err = app.migrator.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "app start error")
	}

	err = app.seeder.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "app start error")
	}

	err = app.svc.Start(ctx)
	if err != nil {
		return errors.Wrap(err, "app start error")
	}

	app.supervisor.AddTasks(
		app.http.Start,
		//app.grpc.Start,
	)

	app.Log().Infof("%s started!", app.Name())

	return app.supervisor.Wait()
}

func (app *App) Stop(ctx context.Context) error {
	return nil
}

func (app *App) Shutdown(ctx context.Context) error {
	return nil
}

func (app *App) EnableSupervisor() {
	name := fmt.Sprintf("%s-supervisor", app.Name())
	app.supervisor = sys.NewSupervisor(name, true, app.opts)
}

func (app *App) SetupRoutes(ctx context.Context) {
	app.http.Setup(ctx)
}

// SetupProbes WIP implementation
func (app *App) SetupProbes(ctx context.Context) {
	health := http2.NewRouter("health", app.opts...)
	health.Mount("/", http2.Healthz)
	app.http.Router().Mount("/healthz", health)
}

func (app *App) RegisterRouter(path string, r http2.Router) {
	app.Log().Infof("Registering '%s' router to handle '%s' routes", r.Name())
	app.http.Router().Mount(path, r)
	app.Log().Infof("'%s' registered", r.Name())
}
