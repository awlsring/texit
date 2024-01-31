package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

func (s *Service) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	return s.apiGw.GetExecution(ctx, id)
}
