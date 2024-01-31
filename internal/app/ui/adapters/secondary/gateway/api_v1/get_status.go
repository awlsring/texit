package apiv1

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	v1 "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (g *ApiGateway) GetNodeStatus(ctx context.Context, id node.Identifier) (node.Status, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.GetNodeStatusRequest{
		Id: id.String(),
	}
	resp, err := g.client.GetNodeStatus(ctx, req)
	if err != nil {
		return node.StatusUnknown, err
	}

	status := TranslateNodeStatus(resp.Status)

	return status, nil
}
