package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) DescribeTailnet(ctx context.Context, req texit.DescribeTailnetParams) (texit.DescribeTailnetRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get tailnet request")

	identifier, err := tailnet.IdentifierFromString(req.Name)
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
	return &texit.DescribeTailnetResponseContent{
		Summary: conversion.TailnetToSummary(tn),
	}, nil
}
