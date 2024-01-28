package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) ListNodes(ctx context.Context) ([]*node.Node, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.ListNodesRequest{}
	resp, err := g.client.ListNodes(ctx, req)
	if err != nil {
		return nil, err
	}

	nodes := []*node.Node{}
	for _, n := range resp.Nodes {
		node, err := SummaryToNode(n)
		if err != nil {
			return nil, err
		}
		nodes = append(nodes, node)
	}

	return nodes, nil
}
