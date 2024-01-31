package interfaces

import (
	"context"

	"github.com/tailscale/tailscale-client-go/tailscale"
)

type Tailscale interface {
	CreateKey(ctx context.Context, caps tailscale.KeyCapabilities, opts ...tailscale.CreateKeyOption) (tailscale.Key, error)
	SetDeviceTags(ctx context.Context, deviceID string, tags []string) error
	DeleteKey(context.Context, string) error
	Devices(ctx context.Context) ([]tailscale.Device, error)
	DeleteDevice(context.Context, string) error
	ACL(ctx context.Context) (*tailscale.ACL, error)
	ValidateACL(ctx context.Context, acl tailscale.ACL) error
	SetACL(ctx context.Context, acl tailscale.ACL, opts ...tailscale.SetACLOption) error
}
