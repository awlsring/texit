package platform_linode

import (
	"context"
	"fmt"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/platform"
	"github.com/google/uuid"
	"github.com/linode/linodego"
)

const (
	DefaultImage = "linode/debian12"
)

func (p *PlatformLinode) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, loc provider.Location, tcs tailnet.ControlServer, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating Linode node")

	log.Debug().Msg("Creating Linode stackscript")
	stack, err := p.client.CreateStackscript(ctx, linodego.StackscriptCreateOptions{
		Label:    id.String(),
		Images:   []string{"any/all"},
		IsPublic: false,
		Script:   platform.TailscaleCloudInit(key.String(), tid.String(), tcs.String()),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Linode stackscript")
		return "", err
	}

	log.Debug().Msg("Creating Linode instance")
	resp, err := p.client.CreateInstance(ctx, linodego.InstanceCreateOptions{
		Region:        loc.String(),
		Type:          DefaultInstanceType,
		Label:         id.String(),
		Tags:          []string{"texit", fmt.Sprintf("tailnet:%s", tid.String()), fmt.Sprintf("node-id:%s", id.String())},
		StackScriptID: stack.ID,
		Image:         DefaultImage,
		RootPass:      uuid.New().String(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Linode instance")
		return "", err
	}

	log.Debug().Msgf("Linode instance created with id %d", resp.ID)
	return node.PlatformIdentifier(fmt.Sprintf("%d", resp.ID)), nil
}
