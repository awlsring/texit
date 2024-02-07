package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) DescribeProvider(ctx context.Context, identifier provider.Identifier) (*provider.Provider, error) {
	req := texit.DescribeProviderParams{
		Name: identifier.String(),
	}
	resp, err := g.client.DescribeProvider(ctx, req)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	switch resp.(type) {
	case *texit.DescribeProviderResponseContent:
		prov, err := SummaryToProvider(resp.(*texit.DescribeProviderResponseContent).Summary)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		return prov, nil
	case *texit.ResourceNotFoundErrorResponseContent:
		return nil, errors.Wrap(gateway.ErrResourceNotFoundError, resp.(*texit.ResourceNotFoundErrorResponseContent).Message)
	case *texit.InvalidInputErrorResponseContent:
		return nil, errors.Wrap(gateway.ErrInvalidInputError, resp.(*texit.InvalidInputErrorResponseContent).Message)
	default:
		return nil, gateway.ErrInternalServerError
	}
}
