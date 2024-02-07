package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) ProvisionNode(ctx context.Context, prov provider.Identifier, loc provider.Location, tn tailnet.Identifier, eph bool) (workflow.ExecutionIdentifier, error) {
	req := &texit.ProvisionNodeRequestContent{
		Provider:  prov.String(),
		Location:  loc.String(),
		Tailnet:   tn.String(),
		Ephemeral: texit.OptBool{},
	}

	resp, err := g.client.ProvisionNode(ctx, req)
	if err != nil {
		return "", translateError(err)
	}

	id, err := workflow.ExecutionIdentifierFromString(resp.(*texit.ProvisionNodeResponseContent).Execution)
	if err != nil {
		return "", translateError(err)
	}

	return id, nil
}
