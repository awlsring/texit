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

func TestStartNode(t *testing.T) {
	ctx := context.Background()
	req := texit.StartNodeParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockNodeSvc.EXPECT().Start(ctx, nodeId).Return(nil)

	res, err := h.StartNode(ctx, req)

	assert.NoError(t, err)
	assert.True(t, res.(*texit.StartNodeResponseContent).Success)
}

func TestStartNodeParseError(t *testing.T) {
	ctx := context.Background()
	req := texit.StartNodeParams{
		Identifier: "",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil)

	res, err := h.StartNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestStartNodeError(t *testing.T) {
	ctx := context.Background()
	req := texit.StartNodeParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockNodeSvc.EXPECT().Start(ctx, nodeId).Return(errors.New("test error"))

	res, err := h.StartNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
