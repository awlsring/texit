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
	ErrUnknownTailnetDevice = errors.New("unknown tailnet device")
)

type Activity interface {
	CreateNodeRecord(context.Context, node.Identifier, node.PlatformIdentifier, provider.Identifier, provider.Location, tailnet.PreauthKey, tailnet.Identifier, tailnet.DeviceIdentifier, tailnet.DeviceName, node.Size, bool) error
	CreateNode(context.Context, provider.Identifier, tailnet.ControlServer, node.Identifier, tailnet.DeviceName, provider.Location, tailnet.PreauthKey, node.Size) (node.PlatformIdentifier, error)
	CreatePreauthKey(context.Context, tailnet.Identifier, bool) (tailnet.PreauthKey, error)
	EnableExitNode(context.Context, tailnet.Identifier, tailnet.DeviceIdentifier) error
	GetDeviceId(context.Context, tailnet.Identifier, tailnet.DeviceName) (tailnet.DeviceIdentifier, error)
	DeleteNode(context.Context, node.Identifier) error
	RemoveTailnetDevice(context.Context, tailnet.Identifier, tailnet.DeviceIdentifier) error
	DeleteNodeRecord(context.Context, node.Identifier) error
	CloseExecution(context.Context, workflow.ExecutionIdentifier, workflow.Status, workflow.ExecutionResult) error
}
