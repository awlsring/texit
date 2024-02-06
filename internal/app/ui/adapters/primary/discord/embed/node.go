package embed

import (
	"fmt"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/app/ui/core/domain/node"
	cnode "github.com/awlsring/texit/internal/pkg/domain/node"
)

func statusToColor(status cnode.Status) uint32 {
	switch status {
	case cnode.StatusRunning:
		return 0x00ff00
	case cnode.StatusStarting:
		return 0xff0000
	case cnode.StatusStopping:
		return 0xff0000
	case cnode.StatusStopped:
		return 0xff0000
	case cnode.StatusUnknown:
		return 0xffff00
	default:
		return 0x000000
	}
}

func NodeAsEmbed(n *node.Node) *tempest.Embed {
	embedTitle := fmt.Sprintf("Node `%s` (`%s`)", n.Identifier.String(), n.TailnetName.String())

	providerField := &tempest.EmbedField{
		Name:   "Provider",
		Value:  fmt.Sprintf("%s (%s)", n.Provider.String(), n.ProviderType.String()),
		Inline: false,
	}
	tailnetField := &tempest.EmbedField{
		Name:   "Tailnet",
		Value:  fmt.Sprintf("%s (%s)", n.Tailnet.String(), n.TailnetType.String()),
		Inline: false,
	}
	locationField := &tempest.EmbedField{
		Name:   "Location",
		Value:  n.Location.String(),
		Inline: true,
	}
	statusField := &tempest.EmbedField{
		Name:   "Status",
		Value:  n.Status.String(),
		Inline: true,
	}
	ephemeralField := &tempest.EmbedField{
		Name:   "Ephemeral",
		Value:  fmt.Sprintf("%t", n.Ephemeral),
		Inline: true,
	}

	return &tempest.Embed{
		Title: embedTitle,
		Color: statusToColor(n.Status),
		Fields: []*tempest.EmbedField{
			providerField,
			tailnetField,
			locationField,
			statusField,
			ephemeralField,
		},
	}
}
