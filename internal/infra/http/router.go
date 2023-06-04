package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/foorester/cook/internal/sys"
)

type (
	Router interface {
		sys.Worker
		chi.Router
		Registerable
	}

	Registerable interface {
		Register(r Registry, path string)
	}

	SimpleRouter struct {
		sys.Worker
		chi.Router
	}
)

func NewRouter(name string, opts ...sys.Option) Router {
	return &SimpleRouter{
		Worker: sys.NewWorker(name, opts...),
		Router: chi.NewRouter(),
	}
}

func (sr *SimpleRouter) Register(r Registry, path string) {
	r.RegisterRouter(path, sr)
}
