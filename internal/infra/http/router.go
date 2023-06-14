package http

import (
	"github.com/go-chi/chi/v5"

	"github.com/foorester/cook/internal/sys"
)

type (
	Router interface {
		sys.Core
		chi.Router
		Registerable
	}

	Registerable interface {
		Register(r Registry, path string)
	}

	SimpleRouter struct {
		sys.Core
		chi.Router
	}
)

func NewRouter(name string, opts ...sys.Option) Router {
	return &SimpleRouter{
		Core:   sys.NewCore(name, opts...),
		Router: chi.NewRouter(),
	}
}

func (sr *SimpleRouter) Register(r Registry, path string) {
	r.RegisterRouter(path, sr)
}
