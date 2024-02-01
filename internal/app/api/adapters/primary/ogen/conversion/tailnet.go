package conversion

import (
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func TranslateTailnetType(t tailnet.Type) texit.TailnetType {
	switch t {
	case tailnet.TypeTailscale:
		return texit.TailnetTypeTailscale
	case tailnet.TypeHeadscale:
		return texit.TailnetTypeHeadscale
	default:
		return texit.TailnetTypeUnknown
	}
}

func TailnetToSummary(t *tailnet.Tailnet) texit.TailnetSummary {
	return texit.TailnetSummary{
		Name:          t.Name.String(),
		Type:          TranslateTailnetType(t.Type),
		ControlServer: t.ControlServer.String(),
	}
}
