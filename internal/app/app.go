package app

import (
	"context"
	"fmt"
	"sync"

	"github.com/foorester/cook/internal/core/service"
	"github.com/foorester/cook/internal/infra/db/pg"
	http2 "github.com/foorester/cook/internal/infra/http"
	pgr "github.com/foorester/cook/internal/infra/repo/pg"
	"github.com/foorester/cook/internal/sys"
	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/errors"
	"github.com/foorester/cook/internal/sys/log"
)

type App struct {
	sync.Mutex
	sys.Worker
	opts       []sys.Option
	supervisor sys.Supervisor
	http       *http2.Server
}

func NewApp(name, namespace string, log log.Logger) (app *App) {
	cfg := config.Load(namespace)

	opts := []sys.Option{
		sys.WithConfig(cfg),
		sys.WithLogger(log),
	}

	app = &App{
		Worker: sys.NewWorker(name, opts...),
		opts:   opts,
	}

	return app
}

func (app *App) Run() (err error) {
	ctx := context.Background()

	err = app.Setup(ctx)
	if err != nil {
		return errors.Wrap(runError, err)
	}

	return app.Start(ctx)
}

func (app *App) Setup(ctx context.Context) error {
	app.EnableSupervisor()

	// Databases
	database := pg.NewDB(app.opts...)

	// Repos
	repo := pgr.NewRecipeRepo(database, app.opts...)

	// Services
	svc := service.NewService(repo, app.opts...)

	// HTTP Server
	app.http = http2.NewServer(app.opts...)
	app.EnableProbes()

	// gRPC servers

	// Event bus

	// WIP: to avoid unused var message
	app.Log().Debugf("Repo: %v", repo)
	app.Log().Debugf("Service: %v", svc)

	return nil
}

func (app *App) Start(ctx context.Context) error {
	app.Log().Infof("%s starting...", app.Name())
	defer app.Log().Infof("%s stopped", app.Name())

	err := app.http.Setup(ctx)
	if err != nil {
		err = errors.Wrap("app start error", err)
		app.Log().Error(err.Error())
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

func (app *App) EnableProbes() {
	health := http2.NewRouter("health", app.opts...)
	health.Mount("/", http2.Healthz)
	app.http.Router().Mount("/healthz", health)
}

func (app *App) RegisterRouter(path string, r http2.Router) {
	app.Log().Infof("Registering '%s' router to handle '%s' routes", r.Name())
	app.http.Router().Mount(path, r)
	app.Log().Infof("'%s' registered", r.Name())
}
