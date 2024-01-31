package conversion

import (
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
)

func TranslateTailnet(t tailnet.Type) teen.Tailnet {
	switch t {
	case tailnet.TypeTailscale:
		return teen.Tailnet_TAILNET_TAILSCALE
	case tailnet.TypeHeadscale:
		return teen.Tailnet_TAILNET_HEADSCALE
	default:
		return teen.Tailnet_TAILNET_UNKNOWN_UNSPECIFIED
	}
}

func TailnetToSummary(t *tailnet.Tailnet) *teen.TailnetSummary {
	return &teen.TailnetSummary{
		Tailnet: t.Name.String(),
		Type:    TranslateTailnet(t.Type),
	}
}
