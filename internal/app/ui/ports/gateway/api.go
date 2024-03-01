package gateway

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/notification"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

var (
	ErrInvalidInputError     = errors.New("invalid input")
	ErrResourceNotFoundError = errors.New("resource not found")
	ErrUnauthorizedError     = errors.New("unauthorized")
	ErrInternalServerError   = errors.New("internal server error")
)

type Api interface {
	DescribeNode(context.Context, node.Identifier) (*node.Node, error)
	GetNodeStatus(context.Context, node.Identifier) (node.Status, error)
	ListNodes(context.Context) ([]*node.Node, error)
	ProvisionNode(context.Context, provider.Identifier, provider.Location, tailnet.Identifier, bool) (workflow.ExecutionIdentifier, error)
	DeprovisionNode(context.Context, node.Identifier) (workflow.ExecutionIdentifier, error)
	StartNode(context.Context, node.Identifier) error
	StopNode(context.Context, node.Identifier) error

	GetExecution(context.Context, workflow.ExecutionIdentifier) (*workflow.Execution, error)

	DescribeProvider(ctx context.Context, id provider.Identifier) (*provider.Provider, error)
	ListProviders(ctx context.Context) ([]*provider.Provider, error)

	DescribeTailnet(ctx context.Context, id tailnet.Identifier) (*tailnet.Tailnet, error)
	ListTailnets(ctx context.Context) ([]*tailnet.Tailnet, error)

	ListNotifiers(ctx context.Context) ([]*notification.Notifier, error)

	HealthCheck(ctx context.Context) error
}
