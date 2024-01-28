package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) StopNode(ctx context.Context, id node.Identifier) error {
	ctx = g.setAuthInContext(ctx)
	req := &v1.StopNodeRequest{
		Id: id.String(),
	}

	_, err := g.client.StopNode(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
