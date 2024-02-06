package provider

import (
	"sync"
	"time"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

type Service struct {
	apiGw       gateway.Api
	mut         sync.Mutex
	lastRefresh time.Time
	providers   map[string]*provider.Provider
}

func NewService(api gateway.Api) service.Provider {
	providers := make(map[string]*provider.Provider)
	return &Service{
		apiGw:     api,
		providers: providers,
	}
}
