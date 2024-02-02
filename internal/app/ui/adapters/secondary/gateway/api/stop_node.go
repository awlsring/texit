package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) StopNode(ctx context.Context, id node.Identifier) error {
	req := texit.StopNodeParams{
		Identifier: id.String(),
	}

	_, err := g.client.StopNode(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
