package handler

import (
	"github.com/awlsring/texit/internal/app/ui/ports/service"
)

type Handler struct {
	apiSvc  service.Api
	nodeSvc service.Node
	provSvc service.Provider
	tailSvc service.Tailnet
}

func New(apiSvc service.Api, nodeSvc service.Node, provSvc service.Provider, tailSvc service.Tailnet) *Handler {
	return &Handler{
		apiSvc:  apiSvc,
		nodeSvc: nodeSvc,
		provSvc: provSvc,
		tailSvc: tailSvc,
	}
}
