package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) StartNode(ctx context.Context, id node.Identifier) error {
	req := texit.StartNodeParams{
		Identifier: id.String(),
	}

	_, err := g.client.StartNode(ctx, req)
	if err != nil {
		return translateError(err)
	}

	return nil
}
