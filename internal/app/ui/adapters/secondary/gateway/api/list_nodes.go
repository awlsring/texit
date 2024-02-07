package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) ListNodes(ctx context.Context) ([]*node.Node, error) {
	resp, err := g.client.ListNodes(ctx)
	if err != nil {
		return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	nodes := []*node.Node{}
	for _, n := range resp.Summaries {
		node, err := SummaryToNode(n)
		if err != nil {
			return nil, errors.Wrap(gateway.ErrInternalServerError, err.Error())
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
