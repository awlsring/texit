package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) ListNodes(ctx context.Context) ([]*node.Node, error) {
	resp, err := g.client.ListNodes(ctx)
	if err != nil {
		return nil, err
	}

	nodes := []*node.Node{}
	for _, n := range resp.(*texit.ListNodesResponseContent).Summaries {
		node, err := SummaryToNode(n)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
