package node

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestList(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	s := &Service{repo: mockRepo}

	testNodes := []*node.Node{
		{Identifier: node.Identifier("test-id-1")},
		{Identifier: node.Identifier("test-id-2")},
	}

	mockRepo.On("List", mock.Anything).Return(testNodes, nil)

	result, err := s.List(ctx)

	assert.NoError(t, err)
	assert.Equal(t, testNodes, result)

	mockRepo.AssertCalled(t, "List", ctx)
}

func TestListError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	s := &Service{repo: mockRepo}

	mockRepo.On("List", mock.Anything).Return(nil, errors.New("test error"))

	result, err := s.List(ctx)

	assert.Error(t, err)
	assert.Nil(t, result)

	mockRepo.AssertCalled(t, "List", ctx)
}
