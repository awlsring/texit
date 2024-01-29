package service

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
)

type Api interface {
	GetNode(context.Context, node.Identifier) (*node.Node, error)
	ListNodes(context.Context) ([]*node.Node, error)
	ProvisionNode(context.Context, provider.Identifier, provider.Location, tailnet.Identifier) (workflow.ExecutionIdentifier, error)
	DeprovisionNode(context.Context, node.Identifier) (workflow.ExecutionIdentifier, error)
	StartNode(context.Context, node.Identifier) error
	StopNode(context.Context, node.Identifier) error

	GetExecution(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)

	GetProvider(ctx context.Context, id provider.Identifier) (*provider.Provider, error)
	ListProviders(ctx context.Context) ([]*provider.Provider, error)
	CheckServerHealth(ctx context.Context) error
}
