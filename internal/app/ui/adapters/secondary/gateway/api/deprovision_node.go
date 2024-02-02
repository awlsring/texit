package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) DeprovisionNode(ctx context.Context, id node.Identifier) (workflow.ExecutionIdentifier, error) {
	req := texit.DeprovisionNodeParams{
		Identifier: id.String(),
	}

	resp, err := g.client.DeprovisionNode(ctx, req)
	if err != nil {
		return "", err
	}

	exId, err := workflow.ExecutionIdentifierFromString(resp.(*texit.DeprovisionNodeResponseContent).Execution)
	if err != nil {
		return "", err
	}

	return exId, nil
}
