package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) ProvisionNode(ctx context.Context, req *teen.ProvisionNodeRequest) (*teen.ProvisionNodeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved provision node request")

	log.Debug().Msg("Validating provision node request")
	provId, err := provider.IdentifierFromString(req.GetProviderId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider identifier")
		return nil, err
	}

	tnId, err := tailnet.IdentifierFromString(req.GetTailnetId())
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
	loc, err := provider.LocationFromString(req.GetLocation(), prov.Platform)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider location")
		return nil, err
	}

	log.Debug().Msg("Launching provision node workflow")
	exId, err := h.workSvc.LaunchProvisionNodeWorkflow(ctx, prov.Name, loc, tail.Name, req.GetEphemeral())
	if err != nil {
		log.Error().Err(err).Msg("Failed to launch provision node workflow")
		return nil, err
	}

	log.Debug().Msg("Successfully launched provision node workflow")
	return &teen.ProvisionNodeResponse{
		ExecutionId: exId.String(),
	}, nil
}
