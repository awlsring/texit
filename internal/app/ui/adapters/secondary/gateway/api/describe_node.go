package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) DescribeNode(ctx context.Context, id node.Identifier) (*node.Node, error) {
	req := texit.DescribeNodeParams{
		Identifier: id.String(),
	}
	resp, err := g.client.DescribeNode(ctx, req)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}
	switch resp := resp.(type) {
	case *texit.DescribeNodeResponseContent:
		n, err := SummaryToNode(resp.Summary)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		return n, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return nil, errors.Wrap(gateway.ErrResourceNotFoundError, resp.Message)
	case *texit.InvalidInputErrorResponseContent:
		return nil, errors.Wrap(gateway.ErrInvalidInputError, resp.Message)
	default:
		return nil, gateway.ErrInternalServerError
	}
}
