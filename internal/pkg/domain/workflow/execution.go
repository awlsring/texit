package workflow

import (
	"time"

	"github.com/pkg/errors"
)

var (
	ErrUnknownStatus   = errors.New("unknown job status")
	ErrUnknownWorkflow = errors.New("unknown workflow")
)

type Execution struct {
	Identifier ExecutionIdentifier
	Workflow   WorkflowName
	Status     Status
	Created    time.Time
	Updated    time.Time
	Finished   *time.Time
	Results    SerializedExecutionResult
}

func NewExecution(id ExecutionIdentifier, workflow WorkflowName) *Execution {
	return &Execution{
		Identifier: id,
		Workflow:   workflow,
		Status:     StatusPending,
		Created:    time.Now(),
		Updated:    time.Now(),
	}
}
