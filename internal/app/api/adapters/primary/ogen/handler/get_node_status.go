package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) GetNodeStatus(ctx context.Context, req texit.GetNodeStatusParams) (texit.GetNodeStatusRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get node status request")

	log.Debug().Msg("Validating node id")
	id, err := node.IdentifierFromString(req.Identifier)
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
	return &texit.GetNodeStatusResponseContent{
		Status: conversion.TranslateNodeStatus(status),
	}, nil
}
