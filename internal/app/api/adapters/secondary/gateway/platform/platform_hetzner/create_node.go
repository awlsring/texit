package platform_hetzner

import (
	"context"
	"fmt"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/platform"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

const (
	DefaultUsServerType = "cpx11"
	DefaultEuServerType = "cx11"
	DefaultImage        = "debian-12"
)

func selectServerTypeForLocation(loc provider.Location) string {
	switch loc.String() {
	case "ash", "hil":
		return DefaultUsServerType
	default:
		return DefaultEuServerType
	}
}

func (p *PlatformHetzner) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, loc provider.Location, tcs tailnet.ControlServer, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating server")

	log.Debug().Msg("calling create server")
	res, _, err := p.client.Server.Create(ctx, hcloud.ServerCreateOpts{
		Name: id.String(),
		ServerType: &hcloud.ServerType{
			Name: selectServerTypeForLocation(loc),
		},
		Image: &hcloud.Image{
			Name: DefaultImage,
		},
		Location: &hcloud.Location{
			Name: loc.String(),
		},
		UserData:         platform.TailscaleCloudInit(key.String(), tid.String(), tcs.String()),
		StartAfterCreate: hcloud.Ptr(true),
		PublicNet: &hcloud.ServerCreatePublicNet{
			EnableIPv4: true,
			EnableIPv6: true,
		},
		Labels: map[string]string{
			"created-by":          "texit",
			"node-id":             id.String(),
			"tailnet-device-name": tid.String(),
		},
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create server")
		return "", err
	}

	log.Debug().Msgf("Hetzner server created with id %d", res.Server.ID)
	return node.PlatformIdentifier(fmt.Sprintf("%d", res.Server.ID)), nil
}
