package handler

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

type Handler struct {
	nodeSvc     service.Node
	providerSvc service.Provider
	teen.UnimplementedTailscaleEphemeralExitNodesServiceServer
}

func New(node service.Node, prov service.Provider) teen.TailscaleEphemeralExitNodesServiceServer {
	return &Handler{
		nodeSvc:     node,
		providerSvc: prov,
	}
}
