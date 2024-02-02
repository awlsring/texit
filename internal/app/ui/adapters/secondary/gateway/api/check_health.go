package api_gateway

import (
	"context"
)

func (g *ApiGateway) HealthCheck(ctx context.Context) error {
	_, err := g.client.Health(ctx)
	if err != nil {
		return err
	}

	return nil
}
