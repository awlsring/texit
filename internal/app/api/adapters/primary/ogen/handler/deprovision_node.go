package handler

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) DeprovisionNode(ctx context.Context, req texit.DeprovisionNodeParams) (texit.DeprovisionNodeRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved delete node request")

	log.Debug().Msg("Validating delete node request")
	nodeId, err := node.IdentifierFromString(req.Identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	log.Debug().Msg("Launching deprovision node workflow")
	exId, err := h.workSvc.LaunchDeprovisionNodeWorkflow(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to launch deprovision node workflow")
		return nil, err
	}

	log.Debug().Msg("Successfully started deprovision node")
	return &texit.DeprovisionNodeResponseContent{
		Execution: exId.String(),
	}, nil
}
