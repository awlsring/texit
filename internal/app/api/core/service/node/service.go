package node

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	repo        repository.Node
	platformGws map[string]gateway.Platform
	workSvc     service.Workflow
}

func NewService(repo repository.Node, workSvc service.Workflow, platformGws map[string]gateway.Platform) service.Node {
	return &Service{
		repo:        repo,
		workSvc:     workSvc,
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
