package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (h *Handler) GetTailnet(ctx context.Context, req *teen.GetTailnetRequest) (*teen.GetTailnetResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get tailnet request")

	identifier, err := tailnet.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse tailnet identifier")
		return nil, err
	}

	tn, err := h.tailnetSvc.Describe(ctx, identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe tailnet")
		return nil, err
	}

	log.Debug().Msg("Successfully described tailnet")
	return &teen.GetTailnetResponse{
		Tailnet: conversion.TailnetToSummary(tn),
	}, nil
}
