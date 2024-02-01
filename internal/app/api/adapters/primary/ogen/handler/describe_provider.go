package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) DescribeProvider(ctx context.Context, req texit.DescribeProviderParams) (texit.DescribeProviderRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get provider request")

	name, err := provider.IdentifierFromString(req.Name)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse provider name")
		return nil, err
	}

	provider, err := h.providerSvc.Describe(ctx, name)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe provider")
		return nil, err
	}

	log.Debug().Msg("Successfully described provider")
	return &texit.DescribeProviderResponseContent{
		Summary: conversion.ProviderToSummary(provider),
	}, nil
}
