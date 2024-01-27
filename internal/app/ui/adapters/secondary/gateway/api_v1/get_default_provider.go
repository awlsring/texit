package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) GetDefaultProvider(ctx context.Context) (*provider.Provider, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.GetDefaultProviderRequest{}
	resp, err := g.client.GetDefaultProvider(ctx, req)
	if err != nil {
		return nil, err
	}

	prov, err := SummaryToProvider(resp.Provider)
	if err != nil {
		return nil, err
	}

	return prov, nil
}
