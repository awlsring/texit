package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type CreatePreauthKeyInput struct {
	TailnetName string `json:"tailnetName"`
	Ephemeral   bool   `json:"ephemeral"`
}

func (h *SfnActivityHandler) createPreAuthKeyActivity(ctx context.Context, input *CreatePreauthKeyInput) (tailnet.PreauthKey, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating pre-auth key")

	log.Debug().Msg("Validating data")
	tn, err := tailnet.IdentifierFromString(input.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet")
		return "", err
	}

	log.Debug().Msg("Creating pre-auth key")
	key, err := h.actSvc.CreatePreauthKey(ctx, tn, input.Ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create pre-auth key")
		return "", err
	}

	return key, nil
}
