package gateway

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
)

var (
	ErrUnknownNode = errors.New("unknown node")
)

type Platform interface {
	GetStatus(context.Context, *node.Node) (node.Status, error)
	DeleteNode(context.Context, *node.Node) error
	StartNode(context.Context, *node.Node) error
	StopNode(context.Context, *node.Node) error
	CreateNode(context.Context, node.Identifier, tailnet.DeviceName, provider.Identifier, provider.Location, tailnet.PreauthKey) (node.PlatformIdentifier, error)
}
