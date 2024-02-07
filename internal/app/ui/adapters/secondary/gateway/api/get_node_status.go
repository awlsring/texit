package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) GetNodeStatus(ctx context.Context, id node.Identifier) (node.Status, error) {
	req := texit.GetNodeStatusParams{
		Identifier: id.String(),
	}
	resp, err := g.client.GetNodeStatus(ctx, req)
	if err != nil {
		return node.StatusUnknown, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	switch resp.(type) {
	case *texit.GetNodeStatusResponseContent:
		status := translateNodeStatus(resp.(*texit.GetNodeStatusResponseContent).Status)

		return status, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return node.StatusUnknown, errors.Wrap(gateway.ErrResourceNotFoundError, resp.(*texit.ResourceNotFoundErrorResponseContent).Message)
	case *texit.InvalidInputErrorResponseContent:
		return node.StatusUnknown, errors.Wrap(gateway.ErrInvalidInputError, resp.(*texit.InvalidInputErrorResponseContent).Message)
	default:
		return node.StatusUnknown, gateway.ErrInternalServerError
	}
}
