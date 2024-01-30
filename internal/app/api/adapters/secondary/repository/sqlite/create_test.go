package sqlite_node_repository

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteNodeRepository{db: db}
	r.initTables(ctx)

	testNode := &node.Node{
		Identifier:         node.Identifier("test-id"),
		PlatformIdentifier: node.PlatformIdentifier("test-platform-id"),
		Provider:           provider.Identifier("test-provider"),
		Location:           provider.Location("test-location"),
		PreauthKey:         tailnet.PreauthKey("test-preauth-key"),
		Tailnet:            tailnet.Identifier("test-tailnet"),
		TailnetName:        tailnet.DeviceName("test-tailnet-name"),
		TailnetIdentifier:  tailnet.DeviceIdentifier("test-tailnet-identifier"),
		Ephemeral:          true,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	err = r.Create(ctx, testNode)

	assert.NoError(t, err)

	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM nodes WHERE identifier = ?", testNode.Identifier.String())
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}
