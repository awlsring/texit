package service

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
)

var (
	ErrExecutionNotFound = errors.New("execution not found")
)

type Workflow interface {
	LaunchProvisionNodeWorkflow(context.Context, provider.Identifier, provider.Location, tailnet.Identifier) (workflow.ExecutionIdentifier, error)
	LaunchDeprovisionNodeWorkflow(context.Context, node.Identifier) (workflow.ExecutionIdentifier, error)
	GetExecution(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)
}
