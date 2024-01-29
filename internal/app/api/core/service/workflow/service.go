package workflow

import (
	"context"
	"sync"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/repository"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

type Service struct {
	nodeRepo    repository.Node
	tailnetGws  map[string]gateway.Tailnet
	platformGws map[string]gateway.Platform

	executions map[string]*workflow.Execution
	mu         sync.Mutex
}

func NewService(nodeRepo repository.Node, tails map[string]gateway.Tailnet, platformGws map[string]gateway.Platform) service.Workflow {
	return &Service{
		nodeRepo:    nodeRepo,
		tailnetGws:  tails,
		platformGws: platformGws,
		executions:  make(map[string]*workflow.Execution),
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

func (s *Service) closeWorkflow(ctx context.Context, ex *workflow.Execution, result workflow.Status) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Closing workflow: %s", ex.Identifier.String())

	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	ex.Updated = now
	ex.Finished = &now
	ex.Status = result
}
