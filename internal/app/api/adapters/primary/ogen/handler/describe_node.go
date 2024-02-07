package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) DescribeNode(ctx context.Context, req texit.DescribeNodeParams) (texit.DescribeNodeRes, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get node request")

	nodeId, err := node.IdentifierFromString(req.Identifier)
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	node, err := h.nodeSvc.Describe(ctx, nodeId)
	if err != nil {
		return nil, err
	}

	log.Debug().Msg("Successfully described node")
	return &texit.DescribeNodeResponseContent{
		Summary: conversion.NodeToSummary(node),
	}, nil
}
