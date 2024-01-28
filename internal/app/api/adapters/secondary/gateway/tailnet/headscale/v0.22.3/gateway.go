package headscale_v0_22_3_gateway

import (
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/gateway"
	"github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/headscale/v0.22.3/client/headscale_service"
)

const (
	oneYearExpiration = 365 * 24 * time.Hour
)

type HeadscaleGateway struct {
	user   string
	client headscale_service.ClientService
}

func New(user string, client headscale_service.ClientService) gateway.Tailnet {
	return &HeadscaleGateway{
		user:   user,
		client: client,
	}
}
