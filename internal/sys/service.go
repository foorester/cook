package sys

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"google.golang.org/grpc"

	"github.com/foorester/cook/internal/sys/config"
	"github.com/foorester/cook/internal/sys/log"
)

type (
	Service interface {
		RegisterHTTPHandler(handler http.Handler)
		RegisterGRPCServer(server *grpc.Server)
	}

	BaseService struct {
		*BaseWorker
	}
)

var (
	runes = []rune("abcdefghijklmnopqrstuvwxyz0123456789")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewService(name string, opts ...Option) *BaseService {
	name = WithSuffix(name, 8)
	return &BaseService{
		BaseWorker: NewWorker(name, opts...),
	}
}

func (bs *BaseService) Init(cfg *config.Config, log log.Logger) {
	bs.SetCfg(cfg)
	bs.SetLog(log)
}

func WithSuffix(name string, n int) string {
	suffix := make([]rune, n)
	for i := range suffix {
		suffix[i] = runes[rand.Intn(len(runes))]
	}
	return fmt.Sprintf("%s-%s", name, string(suffix))
}

type IgnoreUnimplementedRegistration struct{}

var _ Service = (*IgnoreUnimplementedRegistration)(nil)

func (IgnoreUnimplementedRegistration) RegisterHTTPHandler(handler http.Handler) {}

func (IgnoreUnimplementedRegistration) RegisterGRPCServer(server *grpc.Server) {}
