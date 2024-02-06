package tailnet

import (
	"sync"
	"time"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

type Service struct {
	apiGw       gateway.Api
	mut         sync.Mutex
	lastRefresh time.Time
	tailnets    map[string]*tailnet.Tailnet
}

func NewService(api gateway.Api) service.Tailnet {
	tailnets := make(map[string]*tailnet.Tailnet)
	return &Service{
		apiGw:    api,
		tailnets: tailnets,
	}
}
