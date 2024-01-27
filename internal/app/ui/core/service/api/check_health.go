package api

import (
	"context"
)

func (s *Service) CheckServerHealth(ctx context.Context) error {
	return s.apiGw.HealthCheck(ctx)
}
