package api_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/go-faster/errors"
)

func (g *ApiGateway) HealthCheck(ctx context.Context) error {
	_, err := g.client.Health(ctx)
	if err != nil {
		return errors.Wrap(gateway.ErrInternalServerError, err.Error())
	}

	return nil
}
