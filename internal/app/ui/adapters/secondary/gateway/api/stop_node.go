package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) StopNode(ctx context.Context, id node.Identifier) error {
	req := texit.StopNodeParams{
		Identifier: id.String(),
	}

	resp, err := g.client.StopNode(ctx, req)
	if err != nil {
		return errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	switch resp.(type) {
	case *texit.StopNodeResponseContent:
		return nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return errors.Wrap(gateway.ErrResourceNotFoundError, resp.(*texit.ResourceNotFoundErrorResponseContent).Message)
	case *texit.InvalidInputErrorResponseContent:
		return errors.Wrap(gateway.ErrInvalidInputError, resp.(*texit.InvalidInputErrorResponseContent).Message)
	default:
		return gateway.ErrInternalServerError
	}
}
