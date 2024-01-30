package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) GetNodeStatus(ctx context.Context, req *teen.GetNodeStatusRequest) (*teen.GetNodeStatusResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get node status request")

	log.Debug().Msg("Validating node id")
	id, err := node.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	log.Debug().Msg("Getting node status")
	status, err := h.nodeSvc.Status(ctx, id)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get node status")
		return nil, err
	}

	log.Debug().Msg("Successfully got node status")
	return &teen.GetNodeStatusResponse{
		Status: conversion.TranslateNodeStatus(status),
	}, nil
}
