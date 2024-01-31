package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/logger"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (h *Handler) GetProvider(ctx context.Context, req *teen.GetProviderRequest) (*teen.GetProviderResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get provider request")

	identifier, err := provider.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider identifier")
		return nil, err
	}

	provider, err := h.providerSvc.Describe(ctx, identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe provider")
		return nil, err
	}

	log.Debug().Msg("Successfully described provider")
	return &teen.GetProviderResponse{
		Provider: conversion.ProviderToSummary(provider),
	}, nil
}
