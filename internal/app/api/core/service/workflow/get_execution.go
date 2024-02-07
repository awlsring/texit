package workflow

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/app/api/ports/service"
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
		if errors.Is(err, repository.ErrExecutionNotFound) {
			log.Warn().Msgf("Execution not found: %s", id)
			return nil, errors.Wrap(service.ErrExecutionNotFound, id.String())
		}
		log.Error().Err(err).Msg("Failed to get execution from repo")
		return nil, errors.Wrap(err, "failed to get execution")
	}

	log.Debug().Msg("Execution found, returning")
	return exec, nil
}
