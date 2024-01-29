package gateway

import (
	"context"
	"errors"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
)

var (
	ErrUnknownDevice = errors.New("unknown device")
)

type Tailnet interface {
	CreatePreauthKey(context.Context) (tailnet.PreauthKey, error)
	DeletePreauthKey(context.Context, tailnet.PreauthKey) error
	EnableExitNode(context.Context, tailnet.DeviceIdentifier) error
	DeleteDevice(context.Context, tailnet.DeviceIdentifier) error
}
