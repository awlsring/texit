package mem_workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (w *Worker) deprovisionNodeWorkflow(ctx context.Context, input *workflow.DeprovisionNodeInput) (workflow.Status, workflow.ExecutionResult) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deprovisioning node")

	results := workflow.NewDeprovisionNodeExecutionResult()

	nodeId, err := node.IdentifierFromString(input.NodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		results.SetError(err.Error())
		return workflow.StatusFailed, results
	}

	tn, err := tailnet.IdentifierFromString(input.Tailnet)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet")
		results.SetError(err.Error())
		return workflow.StatusFailed, results
	}

	tnId, err := tailnet.DeviceIdentifierFromString(input.TailnetDeviceId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet device id")
		results.SetError(err.Error())
		return workflow.StatusFailed, results
	}

	log.Debug().Msgf("Deleting node")
	err = w.actSvc.DeleteNode(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node")
		results.SetError(err.Error())
		return workflow.StatusFailed, results
	}

	log.Debug().Msg("Removing tailnet device")
	err = w.actSvc.RemoveTailnetDevice(ctx, tn, tnId)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to remove tailnet device")
		results.SetError(err.Error())
		results.ResourcesFailedToDelete = append(results.ResourcesFailedToDelete, "tailnet device")
	}

	log.Debug().Msg("Deleting node record")
	err = w.actSvc.DeleteNodeRecord(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node record")
		results.SetError(err.Error())
		results.ResourcesFailedToDelete = append(results.ResourcesFailedToDelete, "node record")
		return workflow.StatusFailed, results
	}

	return workflow.StatusComplete, results
}
