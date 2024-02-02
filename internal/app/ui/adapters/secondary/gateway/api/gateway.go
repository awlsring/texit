package api_gateway

import (
	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/pkg/gen/texit"
)

type ApiGateway struct {
	client texit.Invoker
}

func New(client texit.Invoker) gateway.Api {
	return &ApiGateway{
		client: client,
	}
}
