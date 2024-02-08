package api_gateway

import (
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func translateExecutionStatus(s texit.ExecutionStatus) workflow.Status {
	switch s {
	case texit.ExecutionStatusRunning:
		return workflow.StatusRunning
	case texit.ExecutionStatusCompleted:
		return workflow.StatusComplete
	case texit.ExecutionStatusFailed:
		return workflow.StatusFailed
	default:
		return workflow.StatusUnknown
	}
}

func translateWorkflowName(s texit.WorkflowName) workflow.WorkflowName {
	switch s {
	case texit.WorkflowNameDeprovisionNode:
		return workflow.WorkflowNameDeprovisionNode
	case texit.WorkflowNameProvisionNode:
		return workflow.WorkflowNameProvisionNode
	default:
		return workflow.WorkflowNameUnknown
	}
}

func float64ToTime(f float64) time.Time {
	return time.Unix(int64(f), 0)
}

func SummaryToExecution(summary texit.ExecutionSummary) (*workflow.Execution, error) {
	eid, err := workflow.ExecutionIdentifierFromString(summary.Identifier)
	if err != nil {
		return nil, err
	}

	var finished time.Time
	if summary.EndedAt.IsSet() {
		finished = float64ToTime(summary.EndedAt.Value)
	}

	ex := &workflow.Execution{
		Identifier: eid,
		Workflow:   translateWorkflowName(summary.Workflow),
		Status:     translateExecutionStatus(summary.Status),
		Created:    float64ToTime(summary.StartedAt),
		Finished:   &finished,
		Results:    workflow.SerializedExecutionResult(summary.Result.Value),
	}

	if summary.EndedAt.IsSet() {
		t := float64ToTime(summary.EndedAt.Value)
		ex.Finished = &t
	}

	return ex, nil
}

func translateNodeStatus(s texit.NodeStatus) node.Status {
	switch s {
	case texit.NodeStatusRunning:
		return node.StatusRunning
	case texit.NodeStatusStarting:
		return node.StatusStarting
	case texit.NodeStatusStopped:
		return node.StatusStopped
	case texit.NodeStatusStopping:
		return node.StatusStopping
	default:
		return node.StatusUnknown
	}
}

func SummaryToNode(summary texit.NodeSummary) (*node.Node, error) {
	nid, err := node.IdentifierFromString(summary.Identifier)
	if err != nil {
		return nil, err
	}

	pid, err := node.PlatformIdentifierFromString(summary.ProviderNodeIdentifier)
	if err != nil {
		return nil, err
	}

	tailId, err := tailnet.IdentifierFromString(summary.Tailnet)
	if err != nil {
		return nil, err
	}

	tailDevId, err := tailnet.DeviceIdentifierFromString(summary.TailnetDeviceIdentifier)
	if err != nil {
		return nil, err
	}

	tailDevName, err := tailnet.DeviceNameFromString(summary.TailnetDeviceName)
	if err != nil {
		return nil, err
	}

	prov, err := provider.IdentifierFromString(summary.Provider)
	if err != nil {
		return nil, err
	}

	location := provider.Location(summary.Location)

	return &node.Node{
		Identifier:         nid,
		Provider:           prov,
		Location:           location,
		PlatformIdentifier: pid,
		Tailnet:            tailId,
		TailnetIdentifier:  tailDevId,
		TailnetName:        tailDevName,
		Ephemeral:          summary.Ephemeral,
		CreatedAt:          float64ToTime(summary.Created),
		UpdatedAt:          float64ToTime(summary.Updated),
	}, nil
}

func translateProviderType(t texit.ProviderType) provider.Type {
	switch t {
	case texit.ProviderTypeAWSEcs:
		return provider.TypeAwsEcs
	case texit.ProviderTypeAWSEc2:
		return provider.TypeAwsEc2
	default:
		return provider.TypeUnknown
	}
}

func SummaryToProvider(summary texit.ProviderSummary) (*provider.Provider, error) {
	name, err := provider.IdentifierFromString(summary.Name)
	if err != nil {
		return nil, err
	}

	return &provider.Provider{
		Name:     name,
		Platform: translateProviderType(summary.Type),
	}, nil
}
func translateTailnetType(t texit.TailnetType) tailnet.Type {
	switch t {
	case texit.TailnetTypeHeadscale:
		return tailnet.TypeHeadscale
	case texit.TailnetTypeTailscale:
		return tailnet.TypeTailscale
	default:
		return tailnet.TypeUnknown
	}
}

func SummaryToTailnet(summary texit.TailnetSummary) (*tailnet.Tailnet, error) {
	id, err := tailnet.IdentifierFromString(summary.Name)
	if err != nil {
		return nil, err
	}

	cs, err := tailnet.ControlServerFromString(summary.ControlServer)
	if err != nil {
		return nil, err
	}

	return &tailnet.Tailnet{
		Name:          id,
		Type:          translateTailnetType(summary.Type),
		ControlServer: cs,
	}, nil
}
