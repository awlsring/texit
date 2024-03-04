package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDescribeNode(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeNodeParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)
	testNode := &node.Node{
		Identifier: nodeId,
	}

	mockNodeSvc.EXPECT().Describe(ctx, nodeId).Return(testNode, nil)

	res, err := h.DescribeNode(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, testNode.Identifier.String(), res.(*texit.DescribeNodeResponseContent).Summary.Identifier)

}

func TestDescribeNodeFailToParse(t *testing.T) {
	ctx := context.Background()

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil, nil)

	badReq := texit.DescribeNodeParams{
		Identifier: "",
	}

	res, err := h.DescribeNode(ctx, badReq)

	assert.Error(t, err)
	assert.Nil(t, res)

}

func TestDescribeNodeError(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeNodeParams{
		Identifier: "test-node",
	}

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil, nil, nil)

	nodeId, _ := node.IdentifierFromString(req.Identifier)

	mockNodeSvc.EXPECT().Describe(mock.Anything, nodeId).Return(nil, errors.New("test error"))

	res, err := h.DescribeNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
