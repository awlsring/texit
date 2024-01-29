package provider

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

type Service struct {
	de      *provider.Provider
	provMap map[provider.Identifier]*provider.Provider
}

func NewService(providers []*provider.Provider) (service.Provider, error) {
	s := &Service{}

	if len(providers) == 0 {
		return nil, service.ErrNoProviders
	}

	provMap := make(map[provider.Identifier]*provider.Provider)
	for _, p := range providers {
		provMap[p.Name] = p
	}

	s.provMap = provMap
	return s, nil
}

func (s *Service) Default(ctx context.Context) (*provider.Provider, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("returning default provider")
	if s.de == nil {
		log.Warn().Msg("no default provider")
		return nil, service.ErrUnknownProvider
	}
	return s.de, nil
}
