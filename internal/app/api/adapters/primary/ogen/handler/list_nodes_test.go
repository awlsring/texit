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

func TestListNodes(t *testing.T) {
	ctx := context.Background()

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil)

	testNodes := []*node.Node{
		{
			Identifier: node.Identifier("test-node-1"),
		},
		{
			Identifier: node.Identifier("test-node-2"),
		},
	}

	mockNodeSvc.EXPECT().List(ctx).Return(testNodes, nil)

	res, err := h.ListNodes(ctx)

	assert.NoError(t, err)
	assert.Len(t, res.(*texit.ListNodesResponseContent).Summaries, len(testNodes))
}

func TestListNodesError(t *testing.T) {
	ctx := context.Background()

	mockNodeSvc := mocks.NewMockNode_service(t)
	h := New(mockNodeSvc, nil, nil, nil)

	mockNodeSvc.EXPECT().List(ctx).Return(nil, errors.New("test error"))

	res, err := h.ListNodes(ctx)

	assert.Error(t, err)
	assert.Nil(t, res)
}
