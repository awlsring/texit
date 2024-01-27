package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) ListProviders(ctx context.Context) ([]*provider.Provider, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.ListProvidersRequest{}
	resp, err := g.client.ListProviders(ctx, req)
	if err != nil {
		return nil, err
	}

	providers := make([]*provider.Provider, len(resp.Providers))
	for i, p := range resp.Providers {
		prov, err := SummaryToProvider(p)
		if err != nil {
			return nil, err
		}

		providers[i] = prov
	}

	return providers, nil
}
