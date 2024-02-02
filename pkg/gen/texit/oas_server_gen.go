// Code generated by ogen, DO NOT EDIT.

package texit

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DeprovisionNode implements DeprovisionNode operation.
	//
	// Deprovision the target node.
	//
	// DELETE /node/{identifier}
	DeprovisionNode(ctx context.Context, params DeprovisionNodeParams) (DeprovisionNodeRes, error)
	// DescribeNode implements DescribeNode operation.
	//
	// Get the summary of an node.
	//
	// GET /node/{identifier}
	DescribeNode(ctx context.Context, params DescribeNodeParams) (DescribeNodeRes, error)
	// DescribeProvider implements DescribeProvider operation.
	//
	// Get the summary of a provider.
	//
	// GET /provider/{name}
	DescribeProvider(ctx context.Context, params DescribeProviderParams) (DescribeProviderRes, error)
	// DescribeTailnet implements DescribeTailnet operation.
	//
	// Get the summary of a provider.
	//
	// GET /tailnet/{name}
	DescribeTailnet(ctx context.Context, params DescribeTailnetParams) (DescribeTailnetRes, error)
	// GetExecution implements GetExecution operation.
	//
	// Get the summary of an execution.
	//
	// GET /execution/{identifier}
	GetExecution(ctx context.Context, params GetExecutionParams) (GetExecutionRes, error)
	// GetNodeStatus implements GetNodeStatus operation.
	//
	// Get the status of an node.
	//
	// GET /node/{identifier}/status
	GetNodeStatus(ctx context.Context, params GetNodeStatusParams) (GetNodeStatusRes, error)
	// Health implements Health operation.
	//
	// GET /health
	Health(ctx context.Context) (*HealthResponseContent, error)
	// ListNodes implements ListNodes operation.
	//
	// Lists all known nodes.
	//
	// GET /node
	ListNodes(ctx context.Context) (ListNodesRes, error)
	// ListProviders implements ListProviders operation.
	//
	// List all registered providers.
	//
	// GET /provider
	ListProviders(ctx context.Context) (*ListProvidersResponseContent, error)
	// ListTailnets implements ListTailnets operation.
	//
	// List all registered tailnets.
	//
	// GET /tailnet
	ListTailnets(ctx context.Context) (*ListTailnetsResponseContent, error)
	// ProvisionNode implements ProvisionNode operation.
	//
	// Provision a node on the specified provider in a given location on the specified tailnet.
	//
	// POST /node
	ProvisionNode(ctx context.Context, req *ProvisionNodeRequestContent) (ProvisionNodeRes, error)
	// StartNode implements StartNode operation.
	//
	// Starts the target node.
	//
	// POST /node/{identifier}/start
	StartNode(ctx context.Context, params StartNodeParams) (StartNodeRes, error)
	// StopNode implements StopNode operation.
	//
	// Stops the target node.
	//
	// POST /node/{identifier}/stop
	StopNode(ctx context.Context, params StopNodeParams) (StopNodeRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}