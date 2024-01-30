package gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
)

type Platform interface {
	DescribeNode(context.Context, *node.Node) (*node.Node, error)
	GetStatus(context.Context, *node.Node) (node.Status, error)
	DeleteNode(context.Context, *node.Node) error
	StartNode(context.Context, *node.Node) error
	StopNode(context.Context, *node.Node) error
	CreateNode(context.Context, node.Identifier, tailnet.DeviceName, provider.Identifier, provider.Location, tailnet.PreauthKey) (node.PlatformIdentifier, error)
}
