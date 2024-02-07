package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) ListProviders(ctx context.Context) ([]*provider.Provider, error) {
	resp, err := g.client.ListProviders(ctx)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	nodes := []*provider.Provider{}
	for _, n := range resp.Summaries {
		node, err := SummaryToProvider(n)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
