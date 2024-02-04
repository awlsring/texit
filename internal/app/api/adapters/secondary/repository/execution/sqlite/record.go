package sqlite_execution_repository

import (
	"strings"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

type ExecutionSqlRecord struct {
	Identifier string     `db:"identifier"`
	Workflow   string     `db:"workflow"`
	Status     string     `db:"status"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  time.Time  `db:"updated_at"`
	FinishedAt *time.Time `db:"finished_at"`
	Results    *string    `db:"results"`
}

func resultsToStringList(results *string) []string {
	if results == nil {
		return nil
	}
	if *results == "" {
		return nil
	}
	return strings.Split(*results, ",")
}

func (n *ExecutionSqlRecord) ToExecution() *workflow.Execution {
	wf, err := workflow.WorkflowNameFromString(n.Workflow)
	if err != nil {
		wf = workflow.WorkflowNameUnknown
	}

	status, err := workflow.JobStatusFromString(n.Status)
	if err != nil {
		status = workflow.StatusUnknown
	}

	return &workflow.Execution{
		Identifier: workflow.ExecutionIdentifier(n.Identifier),
		Workflow:   wf,
		Status:     status,
		Created:    n.CreatedAt,
		Updated:    n.UpdatedAt,
		Finished:   n.FinishedAt,
		Results:    resultsToStringList(n.Results),
	}
}