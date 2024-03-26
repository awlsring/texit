package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type CreateNodeRecordInput struct {
	NodeId            string `json:"nodeId"`
	ProviderName      string `json:"providerName"`
	Location          string `json:"location"`
	TailnetName       string `json:"tailnetName"`
	TailnetDeviceName string `json:"tailnetDeviceName"`
	Size              string `json:"size"`
	Ephemeral         bool   `json:"ephemeral"`
}

func (h *SfnActivityHandler) createNodeRecordActivity(ctx context.Context, input *CreateNodeRecordInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node record")

	log.Debug().Msg("Validating data")
	nodeId, err := node.IdentifierFromString(input.NodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return err
	}

	prov, err := provider.IdentifierFromString(input.ProviderName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		return err
	}

	location, err := provider.LocationFromString(input.Location)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse location")
		return err
	}

	tn, err := tailnet.IdentifierFromString(input.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet name")
		return err
	}

	dev, err := tailnet.DeviceNameFromString(input.TailnetDeviceName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet device name")
		return err
	}

	size, err := node.SizeFromString(input.Size)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node size")
		return err
	}

	log.Debug().Msg("Creating node record")
	_, err = h.actSvc.CreateNodeRecord(ctx, nodeId, prov, location, tn, dev, size, input.Ephemeral)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node record")
		return err
	}

	return nil

}
