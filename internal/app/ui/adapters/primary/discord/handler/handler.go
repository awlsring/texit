package handler

import (
	pending_execution "github.com/awlsring/texit/internal/app/ui/adapters/primary/discord/execution"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
)

type Handler struct {
	apiSvc  service.Api
	nodeSvc service.Node
	provSvc service.Provider
	tailSvc service.Tailnet
	tracker pending_execution.Tracker
}

func New(apiSvc service.Api, nodeSvc service.Node, provSvc service.Provider, tailSvc service.Tailnet, tracker pending_execution.Tracker) *Handler {
	return &Handler{
		apiSvc:  apiSvc,
		nodeSvc: nodeSvc,
		provSvc: provSvc,
		tailSvc: tailSvc,
		tracker: tracker,
	}
}
