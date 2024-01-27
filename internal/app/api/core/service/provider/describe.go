package provider

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Describe(ctx context.Context, name provider.Identifier) (*provider.Provider, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("describing provider %s", name.String())
	p, ok := s.provMap[name]
	if !ok {
		log.Warn().Msg("unknown provider")
		return nil, errors.Wrap(service.ErrUnknownProvider, name.String())
	}

	log.Debug().Msg("returning provider")
	return p, nil
}
