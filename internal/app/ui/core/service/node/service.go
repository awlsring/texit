package node

import (
	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
)

type Service struct {
	apiGw   gateway.Api
	tailSvc service.Tailnet
	provSvc service.Provider
}

func NewService(api gateway.Api, t service.Tailnet, p service.Provider) service.Node {
	return &Service{
		apiGw:   api,
		tailSvc: t,
		provSvc: p,
	}
}
