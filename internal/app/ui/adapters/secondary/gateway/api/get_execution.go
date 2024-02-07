package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	log := logger.FromContext(ctx)
	log.Debug().Str("execution_id", id.String()).Msg("getting execution")

	req := texit.GetExecutionParams{
		Identifier: id.String(),
	}

	log.Debug().Str("execution_id", id.String()).Msg("getting execution")
	resp, err := g.client.GetExecution(ctx, req)
	if err != nil {
		log.Error().Err(err).Msg("failed to get execution")
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	log.Debug().Str("execution_id", id.String()).Msg("got execution, translating response")
	switch resp.(type) {
	case *texit.GetExecutionResponseContent:
		log.Debug().Str("execution_id", id.String()).Msg("response is standard")
		ex, err := SummaryToExecution(resp.(*texit.GetExecutionResponseContent).Summary)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}

		return ex, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		log.Warn().Str("execution_id", id.String()).Msg("response is not found")
		return nil, errors.Wrap(gateway.ErrResourceNotFoundError, resp.(*texit.ResourceNotFoundErrorResponseContent).Message)
	case *texit.InvalidInputErrorResponseContent:
		log.Warn().Str("execution_id", id.String()).Msg("response is invalid input")
		return nil, errors.Wrap(gateway.ErrInvalidInputError, resp.(*texit.InvalidInputErrorResponseContent).Message)
	default:
		log.Error().Str("execution_id", id.String()).Msg("response is unknown, marking as internal server error")
		return nil, gateway.ErrInternalServerError
	}
}
