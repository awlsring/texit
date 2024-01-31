package handler

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func (h *Handler) HealthCheck(ctx context.Context, _ *teen.HealthCheckRequest) (*teen.HealthCheckResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved health check request")
	return &teen.HealthCheckResponse{
		Health: true,
	}, nil
}
