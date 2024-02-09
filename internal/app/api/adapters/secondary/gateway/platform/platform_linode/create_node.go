package platform_linode

import (
	"context"
	"fmt"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/internal/pkg/platform"
	"github.com/google/uuid"
	"github.com/linode/linodego"
)

const (
	DefaultImage    = "linode/debian12"
	PostCreateDelay = 90
)

func (p *PlatformLinode) CreateNode(ctx context.Context, id node.Identifier, tid tailnet.DeviceName, pid *provider.Provider, loc provider.Location, tn *tailnet.Tailnet, key tailnet.PreauthKey) (node.PlatformIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating Linode node")

	log.Debug().Msg("Creating Linode stackscript")
	stack, err := p.client.CreateStackscript(ctx, linodego.StackscriptCreateOptions{
		Label:    id.String(),
		Images:   []string{"any/all"},
		IsPublic: false,
		Script:   platform.TailscaleCloudInit(key.String(), tid.String()),
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
		Tags:          []string{"texit"},
		StackScriptID: stack.ID,
		Image:         DefaultImage,
		RootPass:      uuid.New().String(),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to create Linode instance")
		return "", err
	}

	log.Debug().Msgf("Instance created. Sleeping for %d seconds", PostCreateDelay)
	time.Sleep(PostCreateDelay * time.Second)

	log.Debug().Msgf("Linode instance created with id %d", resp.ID)
	return node.PlatformIdentifier(fmt.Sprintf("%d", resp.ID)), nil
}
