package service

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

type Api interface {
	GetNode(context.Context, node.Identifier) (*node.Node, error)
	GetNodeStatus(context.Context, node.Identifier) (node.Status, error)
	ListNodes(context.Context) ([]*node.Node, error)
	ProvisionNode(context.Context, provider.Identifier, provider.Location, tailnet.Identifier, bool) (workflow.ExecutionIdentifier, error)
	DeprovisionNode(context.Context, node.Identifier) (workflow.ExecutionIdentifier, error)
	StartNode(context.Context, node.Identifier) error
	StopNode(context.Context, node.Identifier) error

	GetExecution(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)

	GetProvider(ctx context.Context, id provider.Identifier) (*provider.Provider, error)
	ListProviders(ctx context.Context) ([]*provider.Provider, error)

	GetTailnet(ctx context.Context, id tailnet.Identifier) (*tailnet.Tailnet, error)
	ListTailnets(ctx context.Context) ([]*tailnet.Tailnet, error)

	CheckServerHealth(ctx context.Context) error
}
