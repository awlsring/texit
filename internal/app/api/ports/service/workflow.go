package service

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

var (
	ErrExecutionNotFound = errors.New("execution not found")
)

type Workflow interface {
	LaunchProvisionNodeWorkflow(context.Context, *provider.Provider, provider.Location, *tailnet.Tailnet, bool) (workflow.ExecutionIdentifier, error)
	LaunchDeprovisionNodeWorkflow(context.Context, node.Identifier) (workflow.ExecutionIdentifier, error)
	GetExecution(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)
}
