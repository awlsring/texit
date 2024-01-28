package apiv1

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) ProvisionNode(ctx context.Context, prov provider.Identifier, loc provider.Location) (workflow.ExecutionIdentifier, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.ProvisionNodeRequest{
		ProviderId: prov.String(),
		Location:   loc.String(),
	}

	resp, err := g.client.ProvisionNode(ctx, req)
	if err != nil {
		return "", err
	}

	id, err := workflow.ExecutionIdentifierFromString(resp.ExecutionId)
	if err != nil {
		return "", err
	}

	return id, nil
}
