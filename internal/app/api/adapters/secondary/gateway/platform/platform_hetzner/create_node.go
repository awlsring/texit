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
	DefaultSmallServerType  = "cpx11"
	DefaultMediumServerType = "cpx21"
	DefaultLargeServerType  = "cpx31"
	DefaultImage            = "debian-12"
)

func instanceSizeForNodeSize(size node.Size) string {
	switch size {
	case node.SizeSmall:
		return DefaultSmallServerType
	case node.SizeMedium:
		return DefaultMediumServerType
	case node.SizeLarge:
		return DefaultLargeServerType
	default:
		return DefaultSmallServerType
	}
}

func (p *PlatformHetzner) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, loc provider.Location, tcs tailnet.ControlServer, key tailnet.PreauthKey, size node.Size) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating server")

	instanceSize := instanceSizeForNodeSize(size)

	log.Debug().Msg("calling create server")
	res, _, err := p.client.Server.Create(ctx, hcloud.ServerCreateOpts{
		Name: id.String(),
		ServerType: &hcloud.ServerType{
			Name: instanceSize,
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
