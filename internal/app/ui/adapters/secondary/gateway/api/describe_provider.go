package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) DescribeProvider(ctx context.Context, identifier provider.Identifier) (*provider.Provider, error) {
	req := texit.DescribeProviderParams{
		Name: identifier.String(),
	}
	resp, err := g.client.DescribeProvider(ctx, req)
	if err != nil {
		return nil, err
	}

	prov, err := SummaryToProvider(resp.(*texit.DescribeProviderResponseContent).Summary)
	if err != nil {
		return nil, err
	}

	return prov, nil
}
