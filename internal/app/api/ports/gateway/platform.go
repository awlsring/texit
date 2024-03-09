package gateway

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

var (
	ErrUnknownNode     = errors.New("unknown node")
	ErrInvalidLocation = errors.New("invalid location")
)

type Platform interface {
	GetStatus(context.Context, *node.Node) (node.Status, error)
	DeleteNode(context.Context, *node.Node) error
	StartNode(context.Context, *node.Node) error
	StopNode(context.Context, *node.Node) error
	CreateNode(context.Context, node.Identifier, tailnet.DeviceName, provider.Location, tailnet.ControlServer, tailnet.PreauthKey, node.Size) (node.PlatformIdentifier, error)
}
