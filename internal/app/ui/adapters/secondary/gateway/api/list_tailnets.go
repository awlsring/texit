package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

func (g *ApiGateway) ListTailnets(ctx context.Context) ([]*tailnet.Tailnet, error) {
	resp, err := g.client.ListTailnets(ctx)
	if err != nil {
		return nil, err
	}

	nodes := []*tailnet.Tailnet{}
	for _, n := range resp.Summaries {
		node, err := SummaryToTailnet(n)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
