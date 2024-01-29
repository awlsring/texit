package tailscale_gateway

import (
	"context"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

const (
	tagCloudExitNode = "tag:cloud-exit-node"
)

type TailscaleGateway struct {
	user   string
	client *tailscale.Client
}

func New(user string, client *tailscale.Client) gateway.Tailnet {
	return &TailscaleGateway{
		user:   user,
		client: client,
	}
}

func (g *TailscaleGateway) CreatePreauthKey(ctx context.Context) (tailnet.PreauthKey, error) {
	log := logger.FromContext(ctx)

	caps := tailscale.KeyCapabilities{
		Devices: struct {
			Create struct {
				Reusable      bool     `json:"reusable"`
				Ephemeral     bool     `json:"ephemeral"`
				Tags          []string `json:"tags"`
				Preauthorized bool     `json:"preauthorized"`
			} `json:"create"`
		}{
			Create: struct {
				Reusable      bool     `json:"reusable"`
				Ephemeral     bool     `json:"ephemeral"`
				Tags          []string `json:"tags"`
				Preauthorized bool     `json:"preauthorized"`
			}{
				Reusable:      false,
				Ephemeral:     false,
				Preauthorized: true,
			},
		},
	}

	log.Info().Msg("creating preauth key")
	resp, err := g.client.CreateKey(ctx, caps)
	if err != nil {
		log.Error().Err(err).Msg("failed to create preauth key")
		return "", err
	}

	log.Debug().Msg("preauth key created")
	return tailnet.PreauthKey(resp.Key), nil
}

func (g *TailscaleGateway) EnableExitNode(ctx context.Context, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Info().Msgf("enabling exit node for %s", tid.String())

	log.Debug().Msg("getting device id")
	id, err := g.findDeviceId(ctx, tid)
	if err != nil {
		log.Error().Err(err).Msg("failed to get device id")
		return err
	}

	log.Debug().Msg("updating acl")
	err = g.updateAcl(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to update acl")
		return err
	}

	log.Debug().Msg("setting device tags")
	err = g.client.SetDeviceTags(ctx, id, []string{tagCloudExitNode})
	if err != nil {
		log.Error().Err(err).Msg("failed to enable exit node")
		return err
	}

	log.Debug().Msg("exit node enabled")
	return nil
}

func (g *TailscaleGateway) DeletePreauthKey(ctx context.Context, key tailnet.PreauthKey) error {
	log := logger.FromContext(ctx)

	log.Info().Msg("deleting preauth key")
	err := g.client.DeleteKey(ctx, key.String())
	if err != nil {
		log.Error().Err(err).Msg("failed to delete preauth key")
		return err
	}

	log.Debug().Msg("preauth key deleted")
	return nil
}

func (g *TailscaleGateway) findDeviceId(ctx context.Context, tid tailnet.DeviceIdentifier) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("finding device id for %s", tid.String())

	for i := 0; i < 3; i++ {
		id, err := g.getDeviceId(ctx, tid)
		if err == nil {
			return id, nil
		}

		log.Warn().Err(err).Msg("failed to get device id, retrying...")
		time.Sleep(time.Second * time.Duration((i+1)*2))
	}

	log.Error().Msgf("failed to get device id for %s", tid.String())
	return "", errors.Wrap(gateway.ErrUnknownDevice, tid.String())
}

func (g *TailscaleGateway) getDeviceId(ctx context.Context, tid tailnet.DeviceIdentifier) (string, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("getting device id for %s", tid.String())

	log.Debug().Msg("listing devices")
	devices, err := g.client.Devices(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to list devices")
		return "", err
	}

	log.Debug().Msg("searching for device")
	for _, device := range devices {
		if device.Hostname == tid.String() {
			log.Debug().Msgf("device found. id: %s", device.ID)
			return device.ID, nil
		}
	}

	log.Error().Msgf("device %s not found", tid.String())
	return "", errors.Wrap(gateway.ErrUnknownDevice, tid.String())
}

func (g *TailscaleGateway) DeleteDevice(ctx context.Context, tid tailnet.DeviceIdentifier) error {
	log := logger.FromContext(ctx)
	log.Info().Msgf("Deleting device %s", tid.String())

	log.Debug().Msg("getting device id")
	id, err := g.findDeviceId(ctx, tid)
	if err != nil {
		log.Error().Err(err).Msg("failed to get device id")
		return err
	}

	log.Debug().Msg("deleting device")
	err = g.client.DeleteDevice(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("failed to delete device")
		return err
	}

	log.Debug().Msg("device deleted")
	return nil
}
