package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	log := logger.FromContext(ctx)
	resp, err := s.apiGw.GetExecution(ctx, id)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			log.Warn().Err(err).Msg("execution not found")
			return nil, service.ErrUnknownExecution
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			log.Warn().Err(err).Msg("invalid input")
			return nil, service.ErrInvalidInputError
		}
		log.Error().Err(err).Msg("failed to get execution")
		return nil, err
	}
	return resp, nil
}
