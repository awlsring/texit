package api

import (
	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
)

// A service that is just a proxy for the Api gateway
type Service struct {
	apiGw gateway.Api
}

func NewService(api gateway.Api) service.Api {
	return &Service{
		apiGw: api,
	}
}
