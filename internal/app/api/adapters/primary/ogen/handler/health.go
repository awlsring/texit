package handler

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func (h *Handler) Health(ctx context.Context) (*texit.HealthResponseContent, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Health check")
	return &texit.HealthResponseContent{
		Healthy: true,
	}, nil
}
