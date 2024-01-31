package apiv1

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (g *ApiGateway) ProvisionNode(ctx context.Context, prov provider.Identifier, loc provider.Location, tn tailnet.Identifier, eph bool) (workflow.ExecutionIdentifier, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.ProvisionNodeRequest{
		ProviderId: prov.String(),
		Location:   loc.String(),
		TailnetId:  tn.String(),
		Ephemeral:  eph,
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
