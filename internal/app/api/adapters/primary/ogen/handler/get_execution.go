package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) GetExecution(ctx context.Context, req texit.GetExecutionParams) (texit.GetExecutionRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get execution request")

	log.Debug().Msg("Validating get execution request")
	exId, err := workflow.ExecutionIdentifierFromString(req.Identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse execution id")
		return nil, err
	}

	log.Debug().Msg("Describing execution")
	ex, err := h.workSvc.GetExecution(ctx, exId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe execution")
		return nil, err
	}

	log.Debug().Msg("Converting execution to summary")
	summary := conversion.ExecutionToSummary(ex)

	log.Debug().Msg("Successfully described execution")
	return &texit.GetExecutionResponseContent{
		Summary: summary,
	}, nil
}
