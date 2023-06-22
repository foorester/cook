package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"golang.org/x/sync/errgroup"

	"github.com/foorester/cook/internal/domain/service"
	"github.com/foorester/cook/internal/infra/openapi"
	"github.com/foorester/cook/internal/sys"
)

type (
	Server struct {
		sys.Core
		opts []sys.Option
		http.Server
		router Router
		svc    service.RecipeService
	}
)

var (
	httpServerName = "http-server"
	apiV1PAth      = "/api/v1"
)

func NewServer(svc service.RecipeService, opts ...sys.Option) (server *Server) {
	return &Server{
		Core:   sys.NewCore(httpServerName, opts...),
		opts:   opts,
		router: NewRouter("root-router", opts...),
		svc:    svc,
	}
}

func (srv *Server) Setup(ctx context.Context) {
	h := NewCookHandler(srv.svc, srv.opts...)

	reqLog := NewReqLoggerMW(srv.Log())

	srv.router.Use(middleware.RequestID)
	srv.router.Use(middleware.RealIP)
	srv.router.Use(reqLog)
	srv.router.Use(middleware.Recoverer)

	srv.router.Mount(apiV1PAth, openapi.Handler(h))
}

func (srv *Server) Start(ctx context.Context) error {
	srv.Server = http.Server{
		Addr:    srv.Address(),
		Handler: srv.Router(),
	}

	group, errGrpCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		srv.Log().Infof("%s started listening at %s", srv.Name(), srv.Address())
		defer srv.Log().Errorf("%s shutdown", srv.Name())

		err := srv.Server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			return err
		}

		return nil
	})

	group.Go(func() error {
		<-errGrpCtx.Done()
		srv.Log().Errorf("%s shutdown", srv.Name())

		ctx, cancel := context.WithTimeout(context.Background(), srv.ShutdownTimeout())
		defer cancel()

		if err := srv.Server.Shutdown(ctx); err != nil {
			return err
		}

		return nil
	})

	return group.Wait()
}

func (srv *Server) SetRouter(r Router) {
	srv.router = r
}

func (srv *Server) Router() (router Router) {
	return srv.router
}

func (srv *Server) Mount(pattern string, handler http.Handler) {
	srv.router.Mount(pattern, handler)
}

func (srv *Server) Address() string {
	host := srv.Cfg().GetString("http.server.host")
	port := srv.Cfg().GetInt("http.server.port")
	return fmt.Sprintf("%s:%d", host, port)
}

func (srv *Server) ShutdownTimeout() time.Duration {
	secs := time.Duration(srv.Cfg().GetInt("http.server.shutdown.timeout.secs"))
	return secs * time.Second
}
