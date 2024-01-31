package apiv1

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	v1 "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (g *ApiGateway) GetNode(ctx context.Context, id node.Identifier) (*node.Node, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.GetNodeRequest{
		Id: id.String(),
	}
	resp, err := g.client.GetNode(ctx, req)
	if err != nil {
		return nil, err
	}

	node, err := SummaryToNode(resp.Node)
	if err != nil {
		return nil, err
	}

	return node, nil
}
