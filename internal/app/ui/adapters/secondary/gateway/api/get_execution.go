package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (g *ApiGateway) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	req := texit.GetExecutionParams{
		Identifier: id.String(),
	}

	resp, err := g.client.GetExecution(ctx, req)
	if err != nil {
		return nil, err
	}

	ex, err := SummaryToExecution(resp.(*texit.GetExecutionResponseContent).Summary)
	if err != nil {
		return nil, err
	}

	return ex, nil
}
