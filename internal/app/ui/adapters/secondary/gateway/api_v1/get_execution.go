package apiv1

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	v1 "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (g *ApiGateway) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	ctx = g.setAuthInContext(ctx)
	req := &v1.GetExecutionRequest{
		ExecutionId: id.String(),
	}

	resp, err := g.client.GetExecution(ctx, req)
	if err != nil {
		return nil, err
	}

	ex, err := SummaryToExecution(resp.Execution)
	if err != nil {
		return nil, err
	}

	return ex, nil
}
