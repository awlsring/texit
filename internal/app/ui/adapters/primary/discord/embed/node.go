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

func statusToEmoji(status cnode.Status) string {
	switch status {
	case cnode.StatusRunning:
		return "🟢"
	case cnode.StatusStarting, cnode.StatusStopping, cnode.StatusStopped:
		return "🔴"
	case cnode.StatusUnknown:
		return "🟡"
	default:
		return "❓"
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
		Value:  fmt.Sprintf("%s %s", statusToEmoji(n.Status), n.Status.String()),
		Inline: true,
	}
	sizeField := &tempest.EmbedField{
		Name:   "Size",
		Value:  n.Size.String(),
		Inline: true,
	}
	ephemeralField := &tempest.EmbedField{
		Name:   "Ephemeral",
		Value:  fmt.Sprintf("%t", n.Ephemeral),
		Inline: true,
	}
	spacerField := &tempest.EmbedField{
		Name:   "\u200b",
		Value:  "\u200b",
		Inline: true,
	}

	return &tempest.Embed{
		Title: embedTitle,
		Color: statusToColor(n.Status),
		Fields: []*tempest.EmbedField{
			providerField,
			tailnetField,
			locationField,
			spacerField,
			statusField,
			sizeField,
			spacerField,
			ephemeralField,
		},
		Footer: &tempest.EmbedFooter{
			Text: fmt.Sprintf("Created at %s", n.CreatedAt.Format("2006-01-02 15:04:05")),
		},
	}
}
