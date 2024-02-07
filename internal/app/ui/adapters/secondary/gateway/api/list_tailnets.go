package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) ListTailnets(ctx context.Context) ([]*tailnet.Tailnet, error) {
	resp, err := g.client.ListTailnets(ctx)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	nodes := []*tailnet.Tailnet{}
	for _, n := range resp.Summaries {
		node, err := SummaryToTailnet(n)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
