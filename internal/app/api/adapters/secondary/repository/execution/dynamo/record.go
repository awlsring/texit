package dynamo_execution_repository

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

const (
	AttributeIdentifier = "identifier"
	AttributeWorkflow   = "workflow"
	AttributeStatus     = "status"
	AttributeCreatedAt  = "created_at"
	AttributeUpdatedAt  = "updated_at"
	AttributeFinishedAt = "finished_at"
	AttributeResults    = "results"
	AttributeTtl        = "ttl"
)

type ExecutionDdbRecord struct {
	Identifier string     `dynamodbav:"identifier"`
	Workflow   string     `dynamodbav:"workflow"`
	Status     string     `dynamodbav:"status"`
	CreatedAt  time.Time  `dynamodbav:"created_at"`
	UpdatedAt  time.Time  `dynamodbav:"updated_at"`
	FinishedAt *time.Time `dynamodbav:"finished_at,omitempty"`
	Results    *string    `dynamodbav:"results,omitempty"`
}

func recordFromExecution(e *workflow.Execution) *ExecutionDdbRecord {
	var results string
	if e.Results != "" {
		results = e.Results.String()
	}
	return &ExecutionDdbRecord{
		Identifier: e.Identifier.String(),
		Workflow:   e.Workflow.String(),
		Status:     e.Status.String(),
		CreatedAt:  e.Created,
		UpdatedAt:  e.Updated,
		FinishedAt: e.Finished,
		Results:    &results,
	}
}

func (n *ExecutionDdbRecord) ToExecution() *workflow.Execution {
	wf, err := workflow.WorkflowNameFromString(n.Workflow)
	if err != nil {
		wf = workflow.WorkflowNameUnknown
	}

	status, err := workflow.StatusFromString(n.Status)
	if err != nil {
		status = workflow.StatusUnknown
	}

	results := workflow.SerializedExecutionResult("")
	if n.Results != nil {
		results = workflow.SerializedExecutionResult(*n.Results)
	}

	return &workflow.Execution{
		Identifier: workflow.ExecutionIdentifier(n.Identifier),
		Workflow:   wf,
		Status:     status,
		Created:    n.CreatedAt,
		Updated:    n.UpdatedAt,
		Finished:   n.FinishedAt,
		Results:    results,
	}
}
