package sqlite_execution_repository

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestGet(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteExecutionRepository{db: db}
	err = r.initTables(ctx)
	assert.NoError(t, err)

	now := time.Now()

	ex := &workflow.Execution{
		Identifier: workflow.FormExecutionIdentifier(workflow.WorkflowNameProvisionNode),
		Workflow:   workflow.WorkflowNameProvisionNode,
		Status:     workflow.StatusRunning,
		Created:    now,
		Updated:    now,
	}

	err = r.CreateExecution(ctx, ex)
	assert.NoError(t, err)

	retrievedNode, err := r.GetExecution(ctx, ex.Identifier)
	assert.NoError(t, err)
	assert.NotNil(t, retrievedNode)
	assert.Equal(t, ex.Identifier, retrievedNode.Identifier)
	assert.Equal(t, ex.Workflow, retrievedNode.Workflow)
	assert.Equal(t, ex.Status, retrievedNode.Status)
}

func TestGet_Closed(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteExecutionRepository{db: db}
	err = r.initTables(ctx)
	assert.NoError(t, err)

	now := time.Now()
	results := []string{"result1", "result2", "result3"}

	ex := &workflow.Execution{
		Identifier: workflow.FormExecutionIdentifier(workflow.WorkflowNameProvisionNode),
		Workflow:   workflow.WorkflowNameProvisionNode,
		Status:     workflow.StatusPending,
		Created:    now,
		Updated:    now,
	}

	err = r.CreateExecution(ctx, ex)
	assert.NoError(t, err)

	err = r.CloseExecution(ctx, ex.Identifier, workflow.StatusComplete, results)
	assert.NoError(t, err)

	retrieved, err := r.GetExecution(ctx, ex.Identifier)
	assert.NoError(t, err)
	assert.NotNil(t, retrieved)
	assert.Equal(t, ex.Identifier, retrieved.Identifier)
	assert.NotNil(t, retrieved.Finished)
	assert.Equal(t, results, retrieved.Results)
}

func TestGet_NotFound(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqliteExecutionRepository{db: db}
	err = r.initTables(ctx)
	assert.NoError(t, err)

	_, err = r.GetExecution(ctx, workflow.ExecutionIdentifier("non-existent-id"))
	assert.Error(t, err)
	assert.True(t, errors.Is(err, repository.ErrNodeNotFound))
}
