package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) StartNode(ctx context.Context, id node.Identifier) error {
	ctx = g.setAuthInContext(ctx)
	req := &v1.StartNodeRequest{
		Id: id.String(),
	}

	_, err := g.client.StartNode(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
