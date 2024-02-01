package conversion

import (
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestTranslateTailnetType(t *testing.T) {
	assert.Equal(t, texit.TailnetTypeTailscale, TranslateTailnetType(tailnet.TypeTailscale))
	assert.Equal(t, texit.TailnetTypeHeadscale, TranslateTailnetType(tailnet.TypeHeadscale))
	assert.Equal(t, texit.TailnetTypeUnknown, TranslateTailnetType(tailnet.TypeUnknown))
}

func TestTailnetToSummary(t *testing.T) {
	tail := &tailnet.Tailnet{
		Name:          tailnet.Identifier("test-name"),
		Type:          tailnet.TypeTailscale,
		ControlServer: tailnet.ControlServer("test-control-server"),
	}

	summary := TailnetToSummary(tail)

	assert.Equal(t, tail.Name.String(), summary.Name)
	assert.Equal(t, TranslateTailnetType(tail.Type), summary.Type)
	assert.Equal(t, tail.ControlServer.String(), summary.ControlServer)
}
