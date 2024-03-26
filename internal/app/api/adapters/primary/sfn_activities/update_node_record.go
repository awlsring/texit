package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type UpdateNodeRecordInput struct {
	NodeId            string `json:"nodeId"`
	PlatformId        string `json:"platformId"`
	ProviderName      string `json:"providerName"`
	Location          string `json:"location"`
	PreauthKey        string `json:"preauthKey"`
	TailnetName       string `json:"tailnetName"`
	TailnetIdentifier string `json:"tailnetDeviceId"`
	TailnetDeviceName string `json:"tailnetDeviceName"`
	Size              string `json:"size"`
	Ephemeral         bool   `json:"ephemeral"`
}

func (h *SfnActivityHandler) updateNodeRecordActivity(ctx context.Context, input *UpdateNodeRecordInput) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Updating node record")

	log.Debug().Msg("Validating data")
	nodeId, err := node.IdentifierFromString(input.NodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return err
	}

	platformId, err := node.PlatformIdentifierFromString(input.PlatformId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse platform id")
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

	key, err := tailnet.PreauthKeyFromString(input.PreauthKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse preauth key")
		return err
	}

	tn, err := tailnet.IdentifierFromString(input.TailnetName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet name")
		return err
	}

	tid, err := tailnet.DeviceIdentifierFromString(input.TailnetIdentifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet identifier")
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

	node := node.Node{
		Identifier:         nodeId,
		PlatformIdentifier: platformId,
		Provider:           prov,
		Tailnet:            tn,
		TailnetIdentifier:  tid,
		TailnetName:        dev,
		Location:           location,
		PreauthKey:         key,
		Size:               size,
		Ephemeral:          input.Ephemeral,
		ProvisionStatus:    node.ProvisionStatusCreated,
	}

	log.Debug().Msg("Updating node record")
	err = h.actSvc.UpdateNodeRecord(ctx, &node)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update node record")
		return err
	}

	return nil

}
