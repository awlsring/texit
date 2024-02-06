package sqlite_execution_repository

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestCloseExecution(t *testing.T) {
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
		Status:     workflow.StatusPending,
		Created:    now,
		Updated:    now,
	}

	err = r.CreateExecution(ctx, ex)
	assert.NoError(t, err)

	err = r.CloseExecution(ctx, ex.Identifier, workflow.StatusComplete, workflow.SerializedExecutionResult(""))
	assert.NoError(t, err)
}
