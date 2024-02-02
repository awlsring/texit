package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
)

func (g *ApiGateway) ListProviders(ctx context.Context) ([]*provider.Provider, error) {
	resp, err := g.client.ListProviders(ctx)
	if err != nil {
		return nil, err
	}

	nodes := []*provider.Provider{}
	for _, n := range resp.Summaries {
		node, err := SummaryToProvider(n)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
