package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) DeprovisionNode(ctx context.Context, id node.Identifier) (workflow.ExecutionIdentifier, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.DeprovisionNodeRequest{
		Id: id.String(),
	}

	resp, err := g.client.DeprovisionNode(ctx, req)
	if err != nil {
		return "", err
	}

	exId, err := workflow.ExecutionIdentifierFromString(resp.ExecutionId)
	if err != nil {
		return "", err
	}

	return exId, nil
}
