package api

import (
	"context"

	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (s *Service) CheckServerHealth(ctx context.Context) error {
	ctx = s.setAuthInContext(ctx)
	_, err := s.client.HealthCheck(ctx, &v1.HealthCheckRequest{})
	if err != nil {
		return err
	}

	return nil
}
