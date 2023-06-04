package http

type (
	Registry interface {
		RegisterRouter(path string, r Router)
	}
)
