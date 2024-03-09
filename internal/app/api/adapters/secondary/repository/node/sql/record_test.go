package sql_node_repository

import (
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/stretchr/testify/assert"
)

func TestToNode(t *testing.T) {
	testRecord := &NodeSqlRecord{
		Identifier:         "test-id",
		PlatformIdentifier: "test-platform-id",
		Provider:           "test-provider",
		Tailnet:            "test-tailnet",
		TailnetName:        "test-tailnet-name",
		TailnetIdentifier:  "test-tailnet-identifier",
		Location:           "test-location",
		PreauthKey:         "test-preauth-key",
		Ephemeral:          true,
		Size:               "small",
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	expectedNode := &node.Node{
		Identifier:         node.Identifier(testRecord.Identifier),
		PlatformIdentifier: node.PlatformIdentifier(testRecord.PlatformIdentifier),
		Provider:           provider.Identifier(testRecord.Provider),
		Tailnet:            tailnet.Identifier(testRecord.Tailnet),
		TailnetIdentifier:  tailnet.DeviceIdentifier(testRecord.TailnetIdentifier),
		TailnetName:        tailnet.DeviceName(testRecord.TailnetName),
		Location:           provider.Location(testRecord.Location),
		PreauthKey:         tailnet.PreauthKey(testRecord.PreauthKey),
		Ephemeral:          testRecord.Ephemeral,
		Size:               node.SizeSmall,
		CreatedAt:          testRecord.CreatedAt,
		UpdatedAt:          testRecord.UpdatedAt,
	}

	resultNode := testRecord.ToNode()

	assert.Equal(t, expectedNode, resultNode)
}
