package workflow

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	nodeRepo    repository.Node
	excRepo     repository.Execution
	tailnetGws  map[string]gateway.Tailnet
	platformGws map[string]gateway.Platform
}

func NewService(nodeRepo repository.Node, excRepo repository.Execution, tails map[string]gateway.Tailnet, platformGws map[string]gateway.Platform) service.Workflow {
	return &Service{
		nodeRepo:    nodeRepo,
		excRepo:     excRepo,
		tailnetGws:  tails,
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

func (s *Service) getTailnetGateway(ctx context.Context, id tailnet.Identifier) (gateway.Tailnet, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting tailnet gateway: %s", id)
	gw, ok := s.tailnetGws[id.String()]
	if !ok {
		log.Error().Msgf("Unknown tailnet: %s", id)
		return nil, errors.Wrap(service.ErrUnknownTailnetId, id.String())
	}
	return gw, nil
}

func (s *Service) closeWorkflow(ctx context.Context, ex workflow.ExecutionIdentifier, result workflow.Status, msgs []string) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Closing workflow: %s", ex.String())

	err := s.excRepo.CloseExecution(ctx, ex, result, msgs)
	if err != nil {
		log.Error().Err(err).Msg("Failed to close workflow")
	}

}
