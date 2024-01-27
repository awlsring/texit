package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (h *Handler) HealthCheck(ctx context.Context, _ *teen.HealthCheckRequest) (*teen.HealthCheckResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved health check request")
	return &teen.HealthCheckResponse{
		Health: true,
	}, nil
}
