package sqlite_node_repository

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteNodeRepository{db: db}
	r.initTables(ctx)

	testNodes := []*node.Node{
		{
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
		},
	}

	for _, n := range testNodes {
		err = r.Create(ctx, n)
		assert.NoError(t, err)
	}

	result, err := r.List(ctx)

	assert.NoError(t, err)
	assert.Len(t, result, len(testNodes))
}

func TestListEmpty(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteNodeRepository{db: db}
	r.initTables(ctx)

	result, err := r.List(ctx)

	assert.NoError(t, err)
	assert.Empty(t, result)
}
