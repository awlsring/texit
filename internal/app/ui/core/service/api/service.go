package api

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/ports/service"
)

type Service struct {
	apiGw gateway.Api
}

func NewService(api gateway.Api) service.Api {
	return &Service{
		apiGw: api,
	}
}
