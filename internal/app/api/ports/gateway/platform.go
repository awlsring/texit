package gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
)

type Platform interface {
	DescribeNode(context.Context, node.PlatformIdentifier) (*node.Node, error)
	DeleteNode(context.Context, node.PlatformIdentifier) error
	StartNode(context.Context, node.PlatformIdentifier) error
	StopNode(context.Context, node.PlatformIdentifier) error
	CreateNode(context.Context, node.Identifier) (*node.Node, error)
}
