package node

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/repository"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	repo        repository.Node
	tailnetGw   gateway.Tailnet
	platformGws map[string]gateway.Platform
}

func NewService(repo repository.Node, tailnetGw gateway.Tailnet, platformGws map[string]gateway.Platform) service.Node {
	return &Service{
		repo:        repo,
		tailnetGw:   tailnetGw,
		platformGws: platformGws,
	}
}

func (s *Service) getPlatformGateway(ctx context.Context, id provider.Identifier) (gateway.Platform, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting platform gateway: %s", id)
	gw, ok := s.platformGws[id.String()]
	if !ok {
		log.Error().Msgf("Unknown platform: %s", id)
		return nil, errors.Wrap(service.ErrUnknownPlatform, id.String())
	}
	return gw, nil
}
