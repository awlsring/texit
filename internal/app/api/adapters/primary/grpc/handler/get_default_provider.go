package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) GetDefaultProvider(ctx context.Context, req *teen.GetDefaultProviderRequest) (*teen.GetDefaultProviderResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get default provider request")

	log.Debug().Msg("Describing default provider")
	provider, err := h.providerSvc.Default(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe provider")
		return nil, err
	}

	log.Debug().Msg("Successfully described default provider")
	return &teen.GetDefaultProviderResponse{
		Provider: conversion.ProviderToSummary(provider),
	}, nil
}
