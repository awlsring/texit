package workflow

import (
	"strings"
	"time"

	"github.com/awlsring/texit/internal/pkg/values"
	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrUnknownStatus   = errors.New("unknown job status")
	ErrUnknownWorkflow = errors.New("unknown workflow")
)

type ExecutionIdentifier string

func (id ExecutionIdentifier) String() string {
	return string(id)
}

func ExecutionIdentifierFromString(id string) (ExecutionIdentifier, error) {
	identifier, err := values.NonNullString[ExecutionIdentifier](id)
	if err != nil {
		return "", err
	}
	return ExecutionIdentifier(identifier), nil
}

func FormExecutionIdentifier(workflow WorkflowName) ExecutionIdentifier {
	id := uuid.New().String()
	return ExecutionIdentifier(workflow.String() + "-" + id)
}

type Status int

const (
	StatusUnknown Status = iota
	StatusPending
	StatusRunning
	StatusComplete
	StatusFailed
)

func (s Status) String() string {
	switch s {
	case StatusUnknown:
		return "unknown"
	case StatusPending:
		return "pending"
	case StatusRunning:
		return "running"
	case StatusComplete:
		return "complete"
	case StatusFailed:
		return "failed"
	default:
		return "unknown"
	}
}

func JobStatusFromString(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "pending":
		return StatusPending, nil
	case "running":
		return StatusRunning, nil
	case "complete":
		return StatusComplete, nil
	case "failed":
		return StatusFailed, nil
	default:
		return StatusUnknown, errors.Wrap(ErrUnknownStatus, s)
	}
}

type WorkflowName int

const (
	WorkflowNameUnknown WorkflowName = iota
	WorkflowNameProvisionNode
	WorkflowNameDeprovisionNode
)

func (n WorkflowName) String() string {
	switch n {
	case WorkflowNameProvisionNode:
		return "provision-node"
	case WorkflowNameDeprovisionNode:
		return "deprovision-node"
	default:
		return "unknown"
	}
}

func WorkflowNameFromString(s string) (WorkflowName, error) {
	switch strings.ToLower(s) {
	case "provision-node":
		return WorkflowNameProvisionNode, nil
	case "deprovision-node":
		return WorkflowNameDeprovisionNode, nil
	default:
		return WorkflowNameUnknown, errors.Wrap(ErrUnknownWorkflow, s)
	}
}

type Execution struct {
	Identifier ExecutionIdentifier
	Workflow   WorkflowName
	Status     Status
	Created    time.Time
	Updated    time.Time
	Finished   *time.Time
	Results    []string
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
