package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) DescribeTailnet(ctx context.Context, identifier tailnet.Identifier) (*tailnet.Tailnet, error) {
	req := texit.DescribeTailnetParams{
		Name: identifier.String(),
	}
	resp, err := g.client.DescribeTailnet(ctx, req)
	if err != nil {
		return nil, translateError(err)
	}

	tail, err := SummaryToTailnet(resp.(*texit.DescribeTailnetResponseContent).Summary)
	if err != nil {
		return nil, translateError(err)
	}

	return tail, nil
}
