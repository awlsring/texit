package handler

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (h *Handler) HealthCheck(ctx context.Context, _ *emptypb.Empty) (*teen.HealthCheckResponse, error) {
	log := logger.FromContext(ctx)
	log.Info().Msg("Recieved health check request")
	return &teen.HealthCheckResponse{
		Health: true,
	}, nil
}
