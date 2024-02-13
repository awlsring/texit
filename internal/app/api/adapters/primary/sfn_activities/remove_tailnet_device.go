package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type RemoveTailnetDeviceInput struct {
	TailnetName     string `json:"tailnetName"`
	TailnetDeviceId string `json:"tailnetDeviceId"`
}

func (h *SfnActivityHandler) removeTailnetDeviceActivity(ctx context.Context, input *RemoveTailnetDeviceInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Removing tailnet device")

	log.Debug().Msg("Validating data")
	tn, err := tailnet.IdentifierFromString(input.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet")
		return err
	}

	tid, err := tailnet.DeviceIdentifierFromString(input.TailnetDeviceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet device id")
		return err
	}

	log.Debug().Msg("Removing tailnet device")
	err = h.actSvc.RemoveTailnetDevice(ctx, tn, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to remove tailnet device")
		return err
	}

	return nil
}
