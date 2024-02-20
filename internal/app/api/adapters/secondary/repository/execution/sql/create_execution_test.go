package sql_execution_repository

import (
	"context"
	"testing"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	_ "modernc.org/sqlite"
)

func TestCreateExecution(t *testing.T) {
	ctx := context.Background()

	db, err := sqlx.Connect("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := &SqlExecutionRepository{db: db}
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
}
