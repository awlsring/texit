package tailscale_gateway

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/interfaces"
)

const (
	tagCloudExitNode = "tag:cloud-exit-node"
)

type TailscaleGateway struct {
	user   string
	client interfaces.Tailscale
}

func New(user string, client interfaces.Tailscale) gateway.Tailnet {
	return &TailscaleGateway{
		user:   user,
		client: client,
	}
}
