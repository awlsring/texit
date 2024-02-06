package workflow

import (
	"github.com/awlsring/texit/internal/pkg/values"
	"github.com/google/uuid"
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
