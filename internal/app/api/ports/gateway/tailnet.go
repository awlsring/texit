package gateway

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

var (
	ErrUnknownDevice = errors.New("unknown device")
)

type Tailnet interface {
	CreatePreauthKey(context.Context, bool) (tailnet.PreauthKey, error)
	DeletePreauthKey(context.Context, tailnet.PreauthKey) error
	EnableExitNode(context.Context, tailnet.DeviceIdentifier) error
	GetDeviceId(context.Context, tailnet.DeviceName) (tailnet.DeviceIdentifier, error)
	DeleteDevice(context.Context, tailnet.DeviceIdentifier) error
}
