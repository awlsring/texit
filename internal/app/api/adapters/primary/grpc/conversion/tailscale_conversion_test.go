package conversion

import (
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
	"github.com/stretchr/testify/assert"
)

func TestTranslateTailnet(t *testing.T) {
	assert.Equal(t, teen.Tailnet_TAILNET_TAILSCALE, TranslateTailnet(tailnet.TypeTailscale))
	assert.Equal(t, teen.Tailnet_TAILNET_HEADSCALE, TranslateTailnet(tailnet.TypeHeadscale))
	assert.Equal(t, teen.Tailnet_TAILNET_UNKNOWN_UNSPECIFIED, TranslateTailnet(tailnet.TypeUnknown))
}

func TestTailnetToSummary(t *testing.T) {
	tail := &tailnet.Tailnet{
		Name: tailnet.Identifier("test-name"),
		Type: tailnet.TypeTailscale,
	}

	summary := TailnetToSummary(tail)

	assert.Equal(t, tail.Name.String(), summary.Tailnet)
	assert.Equal(t, TranslateTailnet(tail.Type), summary.Type)
}
