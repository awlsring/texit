package handler

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func determineSize(size texit.OptNodeSize) node.Size {
	if !size.IsSet() {
		return node.SizeSmall
	}
	switch size.Value {
	case texit.NodeSizeSmall:
		return node.SizeSmall
	case texit.NodeSizeMedium:
		return node.SizeMedium
	case texit.NodeSizeLarge:
		return node.SizeLarge
	default:
		return node.SizeSmall
	}
}

func (h *Handler) ProvisionNode(ctx context.Context, req *texit.ProvisionNodeRequestContent) (texit.ProvisionNodeRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved provision node request")

	log.Debug().Msg("Validating provision node request")
	provId, err := provider.IdentifierFromString(req.Provider)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider identifier")
		return nil, err
	}

	tnId, err := tailnet.IdentifierFromString(req.Tailnet)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet identifier")
		return nil, err
	}

	log.Debug().Msg("Describing provider")
	prov, err := h.providerSvc.Describe(ctx, provId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe provider")
		return nil, err
	}

	log.Debug().Msg("Describing tailnet")
	tail, err := h.tailnetSvc.Describe(ctx, tnId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe tailnet")
		return nil, err
	}

	log.Debug().Msg("Parsing provider location")
	loc, err := provider.LocationFromString(req.GetLocation())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider location")
		return nil, err
	}

	log.Debug().Msg("Parsing node size")
	size := determineSize(req.GetSize())

	log.Debug().Msg("Launching provision node workflow")
	exId, err := h.workSvc.LaunchProvisionNodeWorkflow(ctx, prov, loc, tail, size, req.Ephemeral.Value)
	if err != nil {
		log.Error().Err(err).Msg("Failed to launch provision node workflow")
		return nil, err
	}

	log.Debug().Msg("Successfully launched provision node workflow")
	return &texit.ProvisionNodeResponseContent{
		Execution: exId.String(),
	}, nil
}
