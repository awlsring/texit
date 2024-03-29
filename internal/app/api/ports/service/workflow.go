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
	ErrInvalidLocation   = errors.New("invalid location")
)

type Workflow interface {
	LaunchProvisionNodeWorkflow(context.Context, *provider.Provider, provider.Location, *tailnet.Tailnet, node.Size, bool) (workflow.ExecutionIdentifier, error)
	LaunchDeprovisionNodeWorkflow(context.Context, node.Identifier) (workflow.ExecutionIdentifier, error)
	GetExecution(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)
}
