package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type DeleteNodeInput struct {
	NodeId string `json:"nodeId"`
}

func (h *SfnActivityHandler) deleteNodeActivity(ctx context.Context, input *DeleteNodeInput) (interface{}, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting node request")

	log.Debug().Msg("Parsing node id")
	nodeId, err := node.IdentifierFromString(input.NodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	log.Debug().Msg("Deleting node")
	err = h.actSvc.DeleteNode(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node")
		return nil, err
	}

	return nil, nil
}
