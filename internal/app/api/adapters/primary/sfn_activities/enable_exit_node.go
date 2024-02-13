package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type EnableExitNodeInput struct {
	TailnetName     string `json:"tailnetName"`
	TailnetDeviceId string `json:"tailnetDeviceId"`
}

func (h *SfnActivityHandler) enableExitNodeActivity(ctx context.Context, input *EnableExitNodeInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Enabling exit node")

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

	log.Debug().Msg("Enabling exit node")
	err = h.actSvc.EnableExitNode(ctx, tn, tid)
	if err != nil {
		log.Error().Err(err).Msg("Failed to enable exit node")
		return err
	}

	return nil
}
