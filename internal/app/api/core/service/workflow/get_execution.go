package workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("Getting execution: %s", id)

	log.Debug().Msg("Getting execution from repo")
	exec, err := s.excRepo.GetExecution(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get execution from repo")
		return nil, errors.Wrap(err, "failed to get execution")
	}

	log.Debug().Msg("Execution found, returning")
	return exec, nil
}
