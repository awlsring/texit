package activity

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) getPlatformGateway(ctx context.Context, id provider.Identifier) (gateway.Platform, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting platform gateway: %s", id)
	gw, ok := s.providerGws[id.String()]
	if !ok {
		log.Error().Msg("Platform gateway not found")
		return nil, errors.New("platform gateway not found")
	}
	return gw, nil
}

func (s *Service) getTailnetGateway(ctx context.Context, id tailnet.Identifier) (gateway.Tailnet, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting tailnet gateway: %s", id)
	gw, ok := s.tailnetGws[id.String()]
	if !ok {
		log.Error().Msg("Tailnet gateway not found")
		return nil, errors.New("tailnet gateway not found")
	}
	return gw, nil
}
