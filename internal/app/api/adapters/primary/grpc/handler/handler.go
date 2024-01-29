package handler

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
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
		nodeSvc:     node,
		workSvc:     work,
		providerSvc: prov,
	}
}
