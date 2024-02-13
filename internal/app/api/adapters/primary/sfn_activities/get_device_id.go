package sfn_activities

import (
	"context"
	"errors"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type GetDeviceIdInput struct {
	TailnetName       string `json:"tailnetName"`
	TailnetDeviceName string `json:"tailnetDeviceName"`
}

func (h *SfnActivityHandler) getDeviceIdActivity(ctx context.Context, input *GetDeviceIdInput) (tailnet.DeviceIdentifier, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Getting device id")

	log.Debug().Msg("Validating data")
	tn, err := tailnet.IdentifierFromString(input.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet")
		return "", err
	}

	dev, err := tailnet.DeviceNameFromString(input.TailnetDeviceName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet device name")
		return "", err
	}

	log.Debug().Msg("Getting device id")
	id, err := h.actSvc.GetDeviceId(ctx, tn, dev)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get device id")
		if errors.Is(err, service.ErrUnknownTailnetDevice) {
			log.Warn().Err(err).Msg("Device not found")
			return "", &DeviceNotFoundError{"Device not found"}
		}
		return "", err
	}

	return id, nil
}
