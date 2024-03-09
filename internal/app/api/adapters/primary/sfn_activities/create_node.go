package sfn_activities

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type CreateNodeInput struct {
	ProviderName         string `json:"providerName"`
	Location             string `json:"location"`
	NodeId               string `json:"nodeId"`
	TailnetDeviceName    string `json:"tailnetDeviceName"`
	TailnetControlServer string `json:"tailnetControlServer"`
	PreauthKey           string `json:"preauthKey"`
	Size                 string `json:"size"`
}

func (h *SfnActivityHandler) createNodeActivity(ctx context.Context, input *CreateNodeInput) (interface{}, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Create Node Activity")

	log.Debug().Msg("Validating data")
	prov, err := provider.IdentifierFromString(input.ProviderName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		return nil, err
	}

	location, err := provider.LocationFromString(input.Location)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse location")
		return nil, err
	}

	nodeId, err := node.IdentifierFromString(input.NodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	tcs, err := tailnet.ControlServerFromString(input.TailnetControlServer)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet control server")
		return nil, err
	}

	dev, err := tailnet.DeviceNameFromString(input.TailnetDeviceName)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet device name")
		return nil, err
	}

	key, err := tailnet.PreauthKeyFromString(input.PreauthKey)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse preauth key")
		return nil, err
	}

	size, err := node.SizeFromString(input.Size)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node size")
		return nil, err
	}

	log.Debug().Msg("Creating node")
	node, err := h.actSvc.CreateNode(ctx, prov, tcs, nodeId, dev, location, key, size)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node")
		return nil, err
	}

	return node, nil
}
