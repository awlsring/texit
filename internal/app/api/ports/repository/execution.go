package repository

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

var (
	ErrExecutionNotFound = errors.New("execution not found")
)

type Execution interface {
	Init(ctx context.Context) error
	GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error)
	CreateExecution(ctx context.Context, ex *workflow.Execution) error
	CloseExecution(ctx context.Context, id workflow.ExecutionIdentifier, result workflow.Status, output workflow.SerializedExecutionResult) error
}
