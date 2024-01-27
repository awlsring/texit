package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) GetProvider(ctx context.Context, req *teen.GetProviderRequest) (*teen.GetProviderResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get provider request")

	identifier, err := provider.ProviderFromString(req.GetId())
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
