package sqlite_node_repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteNodeRepository{db: db}
	err = r.initTables(ctx)
	assert.NoError(t, err)

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

	retrievedNode, err := r.Get(ctx, testNode.Identifier)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedNode)
	assert.Equal(t, testNode.Identifier, retrievedNode.Identifier)
}

func TestGet_NotFound(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteNodeRepository{db: db}
	err = r.initTables(ctx)
	assert.NoError(t, err)

	_, err = r.Get(ctx, node.Identifier("non-existent-id"))
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNodeNotFound))
}