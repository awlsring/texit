package conversion

import (
	"testing"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/stretchr/testify/assert"
)

func TestNodeToSummary(t *testing.T) {
	id, err := node.IdentifierFromString("test-id")
	assert.NoError(t, err)

	providerId, err := provider.IdentifierFromString("test-provider")
	assert.NoError(t, err)

	platformId, err := node.PlatformIdentifierFromString("test-platform-id")
	assert.NoError(t, err)

	tailnetId, err := tailnet.DeviceIdentifierFromString("test-tailnet-id")
	assert.NoError(t, err)

	tn, err := tailnet.IdentifierFromString("test-tailnet")
	assert.NoError(t, err)

	loc, err := provider.LocationFromString("us-west-2", provider.TypeAwsEcs)
	assert.NoError(t, err)

	tailnetName := tailnet.FormDeviceName(loc.String(), id.String())

	n := &node.Node{
		Identifier:         id,
		Provider:           providerId,
		PlatformIdentifier: platformId,
		TailnetIdentifier:  tailnetId,
		Tailnet:            tn,
		TailnetName:        tailnetName,
		Location:           loc,
		Ephemeral:          true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	summary := NodeToSummary(n)

	assert.Equal(t, n.Identifier.String(), summary.Id)
	assert.Equal(t, n.Provider.String(), summary.Provider)
	assert.Equal(t, n.PlatformIdentifier.String(), summary.PlatformId)
	assert.Equal(t, n.TailnetIdentifier.String(), summary.TailnetId)
	assert.Equal(t, n.Tailnet.String(), summary.Tailnet)
	assert.Equal(t, n.TailnetName.String(), summary.TailnetName)
	assert.Equal(t, n.Location.String(), summary.Location)
	assert.Equal(t, n.Ephemeral, summary.Ephemeral)
	assert.Equal(t, n.CreatedAt.Format(time.RFC3339Nano), summary.CreatedAt)
	assert.Equal(t, n.UpdatedAt.Format(time.RFC3339Nano), summary.UpdatedAt)
}
