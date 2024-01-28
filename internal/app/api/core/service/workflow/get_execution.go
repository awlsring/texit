package workflow

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting execution: %s", id)

	log.Debug().Msg("Getting execution from local map")
	s.mu.Lock()
	defer s.mu.Unlock()
	exec, ok := s.executions[id.String()]
	if !ok {
		log.Error().Msgf("Unknown execution: %s", id)
		return nil, errors.Wrap(service.ErrExecutionNotFound, id.String())
	}

	log.Debug().Msg("Execution found, returning")
	return exec, nil
}
