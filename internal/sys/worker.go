package sys

import (
	"context"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/log"
)

type (
	Worker interface {
		Name() string
		Log() log.Logger
		Cfg() *config.Config
		Setup(ctx context.Context)
		Start(ctx context.Context) error
		Stop(ctx context.Context) error
	}
)

type (
	BaseWorker struct {
		name     string
		log      log.Logger
		cfg      *config.Config
		didSetup bool
		didStart bool
	}
)

func NewWorker(name string, opts ...Option) *BaseWorker {
	name = GenName(name, "worker")

	bw := &BaseWorker{
		name: name,
	}

	for _, opt := range opts {
		opt(bw)
	}

	return bw
}

func (bw *BaseWorker) Name() string {
	return bw.name
}

func (bw *BaseWorker) SetName(name string) {
	bw.name = name
}

func (bw *BaseWorker) Log() log.Logger {
	return bw.log
}

func (bw *BaseWorker) SetLog(log log.Logger) {
	bw.log = log
}

func (bw *BaseWorker) Cfg() *config.Config {
	return bw.cfg
}

func (bw *BaseWorker) SetCfg(cfg *config.Config) {
	bw.cfg = cfg
}

func (bw *BaseWorker) Setup(ctx context.Context) {
	bw.Log().Infof("%s setup", bw.Name())
}

func (bw *BaseWorker) Start(ctx context.Context) error {
	bw.Log().Infof("%s start", bw.Name())
	return nil
}

func (bw *BaseWorker) Stop(ctx context.Context) error {
	bw.Log().Infof("%s stop", bw.Name())
	return nil
}

func GenName(name, defName string) string {
	if strings.Trim(name, " ") == "" {
		return fmt.Sprintf("%s-%s", defName, nameSufix())
	}
	return name
}

func nameSufix() string {
	digest := hash(time.Now().String())
	return digest[len(digest)-8:]
}

func hash(s string) string {
	h := fnv.New32a()
	h.Write([]byte(s))
	return fmt.Sprintf("%d", h.Sum32())
}

type (
	Option func(w *BaseWorker)
)

func WithConfig(cfg *config.Config) Option {
	return func(svc *BaseWorker) {
		svc.SetCfg(cfg)
	}
}

func WithLogger(log log.Logger) Option {
	return func(svc *BaseWorker) {
		svc.SetLog(log)
	}
}
