package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type DeleteNodeRecordInput struct {
	NodeId string `json:"nodeId"`
}

func (h *SfnActivityHandler) deleteNodeRecordActivity(ctx context.Context, input *DeleteNodeRecordInput) (interface{}, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting node record request")

	log.Debug().Msg("Parsing node id")
	nodeId, err := node.IdentifierFromString(input.NodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	log.Debug().Msg("Deleting node record")
	err = h.actSvc.DeleteNodeRecord(ctx, nodeId)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to delete node record")
		return nil, err
	}

	return nil, nil
}
