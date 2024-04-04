package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/api/adapters/primary/ogen/conversion"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestGetNodeStatus(t *testing.T) {
	ctx := context.Background()
	req := texit.GetNodeStatusParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)
	mockNode := node.Node{
		Identifier:      nodeId,
		ProvisionStatus: node.ProvisionStatusCreated,
	}
	testStatus := node.StatusRunning

	mockNodeSvc.EXPECT().Describe(ctx, nodeId).Return(&mockNode, nil)
	mockNodeSvc.EXPECT().Status(ctx, nodeId).Return(testStatus, nil)

	res, err := h.GetNodeStatus(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, conversion.TranslateNodeStatus(testStatus), res.(*texit.GetNodeStatusResponseContent).Status)
}

func TestGetNodeStatusFailToParse(t *testing.T) {
	ctx := context.Background()

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil, nil)

	badReq := texit.GetNodeStatusParams{
		Identifier: "",
	}

	res, err := h.GetNodeStatus(ctx, badReq)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestGetNodeStatusError(t *testing.T) {
	ctx := context.Background()
	req := texit.GetNodeStatusParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockNodeSvc.EXPECT().Describe(ctx, nodeId).Return(nil, errors.New("test error"))

	res, err := h.GetNodeStatus(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
