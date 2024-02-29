package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestStopNode(t *testing.T) {
	ctx := context.Background()
	req := texit.StopNodeParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockNodeSvc.EXPECT().Stop(ctx, nodeId).Return(nil)

	res, err := h.StopNode(ctx, req)

	assert.NoError(t, err)
	assert.True(t, res.(*texit.StopNodeResponseContent).Success)
}

func TestStopNodeParseError(t *testing.T) {
	ctx := context.Background()
	req := texit.StopNodeParams{
		Identifier: "",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil)

	res, err := h.StopNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestStopNodeError(t *testing.T) {
	ctx := context.Background()
	req := texit.StopNodeParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockNodeSvc.EXPECT().Stop(ctx, nodeId).Return(errors.New("test error"))

	res, err := h.StopNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
