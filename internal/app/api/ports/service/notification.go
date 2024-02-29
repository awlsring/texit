package service

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

// Service that handles notifying callers about completion of requests
type Notification interface {
	ListNotifiers(ctx context.Context) []*notification.Notifier
	NotifyExecutionCompletion(ctx context.Context, e workflow.ExecutionIdentifier, w workflow.WorkflowName, status workflow.Status, results workflow.ExecutionResult) error
}
