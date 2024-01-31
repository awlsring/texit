package apiv1

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	v1 "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (g *ApiGateway) GetProvider(ctx context.Context, identifier provider.Identifier) (*provider.Provider, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.GetProviderRequest{
		Id: identifier.String(),
	}
	resp, err := g.client.GetProvider(ctx, req)
	if err != nil {
		return nil, err
	}

	prov, err := SummaryToProvider(resp.Provider)
	if err != nil {
		return nil, err
	}

	return prov, nil
}
