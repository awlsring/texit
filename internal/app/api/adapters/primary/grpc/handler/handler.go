package handler

import (
	"github.com/awlsring/texit/internal/app/api/ports/service"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

type Handler struct {
	nodeSvc     service.Node
	workSvc     service.Workflow
	providerSvc service.Provider
	tailnetSvc  service.Tailnet
	teen.UnimplementedTailscaleEphemeralExitNodesServiceServer
}

func New(node service.Node, work service.Workflow, prov service.Provider, tail service.Tailnet) teen.TailscaleEphemeralExitNodesServiceServer {
	return &Handler{
		tailnetSvc:  tail,
		nodeSvc:     node,
		workSvc:     work,
		providerSvc: prov,
	}
}
