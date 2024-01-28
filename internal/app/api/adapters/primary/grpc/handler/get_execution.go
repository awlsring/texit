package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/adapters/primary/grpc/conversion"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) GetExecution(ctx context.Context, req *v1.GetExecutionRequest) (*v1.GetExecutionResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved get execution request")

	log.Debug().Msg("Validating get execution request")
	exId, err := workflow.ExecutionIdentifierFromString(req.GetExecutionId())
	if err != nil {
		log.Error().Err(err).Msg("Failed to parse execution id")
		return nil, err
	}

	log.Debug().Msg("Describing execution")
	ex, err := h.workSvc.GetExecution(ctx, exId)
	if err != nil {
		log.Error().Err(err).Msg("Failed to describe execution")
		return nil, err
	}

	log.Debug().Msg("Successfully described execution")
	return &v1.GetExecutionResponse{
		Execution: conversion.ExecutionToSummary(ex),
	}, nil
}
