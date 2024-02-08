package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) DescribeTailnet(ctx context.Context, identifier tailnet.Identifier) (*tailnet.Tailnet, error) {
	req := texit.DescribeTailnetParams{
		Name: identifier.String(),
	}
	resp, err := g.client.DescribeTailnet(ctx, req)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	switch resp := resp.(type) {
	case *texit.DescribeTailnetResponseContent:
		tail, err := SummaryToTailnet(resp.Summary)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}

		return tail, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return nil, errors.Wrap(gateway.ErrResourceNotFoundError, resp.Message)
	case *texit.InvalidInputErrorResponseContent:
		return nil, errors.Wrap(gateway.ErrInvalidInputError, resp.Message)
	default:
		return nil, gateway.ErrInternalServerError
	}
}
