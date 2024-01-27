package gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
)

type Tailnet interface {
	CreatePreauthKey(context.Context) (tailnet.PreauthKey, error)
	DeletePreauthKey(context.Context, tailnet.PreauthKey) error
	DeleteDevice(context.Context, tailnet.DeviceIdentifier) error
}
