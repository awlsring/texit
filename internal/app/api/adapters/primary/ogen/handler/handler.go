package handler

import (
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/pkg/gen/texit"
)

type Handler struct {
	nodeSvc     service.Node
	workSvc     service.Workflow
	providerSvc service.Provider
	notifierSvc service.Notification
	tailnetSvc  service.Tailnet
}

func New(nodeSvc service.Node, workSvc service.Workflow, provSvc service.Provider, tailSvc service.Tailnet, notSvc service.Notification) texit.Handler {
	return &Handler{
		nodeSvc:     nodeSvc,
		workSvc:     workSvc,
		providerSvc: provSvc,
		tailnetSvc:  tailSvc,
		notifierSvc: notSvc,
	}
}
