package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type FormIdentifiersInput struct {
	Location string `json:"location"`
}

type FormIdentifiersOutput struct {
	NodeId            string `json:"nodeId"`
	TailnetDeviceName string `json:"tailnetDeviceName"`
}

func (h *SfnActivityHandler) formIdentifiersActivity(ctx context.Context, input *FormIdentifiersInput) (*FormIdentifiersOutput, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Closing execution request")

	log.Debug().Msg("Validating input")
	location, err := provider.LocationFromString(input.Location)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse location")
		return nil, nil
	}

	log.Debug().Msg("Forming node id")
	id := node.FormNewNodeIdentifier()
	log.Debug().Msgf("New node id: %s", id)

	log.Debug().Msg("Forming tailnet identifier")
	tailName := tailnet.FormDeviceName(location.String(), id.String())
	log.Debug().Msgf("New tailnet device name: %s", tailName)

	return &FormIdentifiersOutput{
		NodeId:            id.String(),
		TailnetDeviceName: tailName.String(),
	}, nil
}
