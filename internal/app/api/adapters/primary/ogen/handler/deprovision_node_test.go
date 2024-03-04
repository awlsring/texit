package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestDeprovisionNode(t *testing.T) {
	ctx := context.Background()
	req := texit.DeprovisionNodeParams{
		Identifier: "test-node",
	}

	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)
	exId := workflow.ExecutionIdentifier("test-execution-id")

	mockWorkSvc.EXPECT().LaunchDeprovisionNodeWorkflow(ctx, nodeId).Return(exId, nil)

	res, err := h.DeprovisionNode(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, exId.String(), res.(*texit.DeprovisionNodeResponseContent).Execution)
}

func TestDeprovisionNodeFailToParse(t *testing.T) {
	ctx := context.Background()

	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, nil, nil, nil, nil)

	badReq := texit.DeprovisionNodeParams{
		Identifier: "",
	}

	res, err := h.DeprovisionNode(ctx, badReq)

	assert.Error(t, err)
	assert.Nil(t, res)

}

func TestDeprovisionNodeError(t *testing.T) {
	ctx := context.Background()
	req := texit.DeprovisionNodeParams{
		Identifier: "test-node",
	}

	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockWorkSvc.EXPECT().LaunchDeprovisionNodeWorkflow(ctx, nodeId).Return("", errors.New("test error"))

	res, err := h.DeprovisionNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
