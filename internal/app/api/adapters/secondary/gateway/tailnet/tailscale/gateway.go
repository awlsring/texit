package tailscale_gateway

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/interfaces"
)

const (
	tagCloudExitNode = "tag:cloud-exit-node"
)

type TailscaleGateway struct {
	client interfaces.Tailscale
}

func New(client interfaces.Tailscale) gateway.Tailnet {
	return &TailscaleGateway{
		client: client,
	}
}
