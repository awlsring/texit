package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) DescribeNode(ctx context.Context, id node.Identifier) (*node.Node, error) {
	req := texit.DescribeNodeParams{
		Identifier: id.String(),
	}
	resp, err := g.client.DescribeNode(ctx, req)
	if err != nil {
		return nil, err
	}

	node, err := SummaryToNode(resp.(*texit.DescribeNodeResponseContent).Summary)
	if err != nil {
		return nil, err
	}

	return node, nil
}
