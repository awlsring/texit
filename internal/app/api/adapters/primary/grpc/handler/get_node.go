package handler

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (h *Handler) GetNode(ctx context.Context, req *teen.GetNodeRequest) (*teen.GetNodeResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get node request")

	nodeId, err := node.IdentifierFromString(req.GetId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse node id")
		return nil, err
	}

	node, err := h.nodeSvc.Describe(ctx, nodeId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe node")
		return nil, err
	}

	log.Debug().Msg("Successfully described node")
	return &teen.GetNodeResponse{
		Node: conversion.NodeToSummary(node),
	}, nil
}
