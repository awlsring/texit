package workflow

import (
	"strings"

	"github.com/pkg/errors"
)

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
