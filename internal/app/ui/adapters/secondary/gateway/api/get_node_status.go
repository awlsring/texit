package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) GetNodeStatus(ctx context.Context, id node.Identifier) (node.Status, error) {
	req := texit.GetNodeStatusParams{
		Identifier: id.String(),
	}
	resp, err := g.client.GetNodeStatus(ctx, req)
	if err != nil {
		return node.StatusUnknown, translateError(err)
	}

	status := translateNodeStatus(resp.(*texit.GetNodeStatusResponseContent).Status)

	return status, nil
}
