package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) StartNode(ctx context.Context, id node.Identifier) error {
	req := texit.StartNodeParams{
		Identifier: id.String(),
	}

	resp, err := g.client.StartNode(ctx, req)
	if err != nil {
		return errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	switch resp := resp.(type) {
	case *texit.StartNodeResponseContent:
		return nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return errors.Wrap(gateway.ErrResourceNotFoundError, resp.Message)
	case *texit.InvalidInputErrorResponseContent:
		return errors.Wrap(gateway.ErrInvalidInputError, resp.Message)
	default:
		return gateway.ErrInternalServerError
	}
}
