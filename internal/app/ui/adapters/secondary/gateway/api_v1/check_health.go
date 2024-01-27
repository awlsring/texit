package apiv1

import (
	"context"

	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func (g *ApiGateway) HealthCheck(ctx context.Context) error {
	ctx = g.setAuthInContext(ctx)
	_, err := g.client.HealthCheck(ctx, &v1.HealthCheckRequest{})
	if err != nil {
		return err
	}

	return nil
}
