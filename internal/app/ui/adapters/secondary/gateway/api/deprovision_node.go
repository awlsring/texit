package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) DeprovisionNode(ctx context.Context, id node.Identifier) (workflow.ExecutionIdentifier, error) {
	req := texit.DeprovisionNodeParams{
		Identifier: id.String(),
	}

	resp, err := g.client.DeprovisionNode(ctx, req)
	if err != nil {
		return "", errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}
	switch resp.(type) {
	case *texit.DeprovisionNodeResponseContent:
		exId, err := workflow.ExecutionIdentifierFromString(resp.(*texit.DeprovisionNodeResponseContent).Execution)
		if err != nil {
			return "", errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		return exId, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return "nil", errors.Wrap(gateway.ErrResourceNotFoundError, resp.(*texit.ResourceNotFoundErrorResponseContent).Message)
	case *texit.InvalidInputErrorResponseContent:
		return "", errors.Wrap(gateway.ErrInvalidInputError, resp.(*texit.InvalidInputErrorResponseContent).Message)
	default:
		return "", gateway.ErrInternalServerError
	}

}
