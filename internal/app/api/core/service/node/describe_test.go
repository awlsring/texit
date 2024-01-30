package node

import (
	"context"
	"testing"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/repository"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDescribe(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	s := &Service{repo: mockRepo}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
	}

	mockRepo.EXPECT().Get(mock.Anything, testNode.Identifier).Return(testNode, nil)

	result, err := s.Describe(ctx, testNode.Identifier)

	assert.NoError(t, err)
	assert.Equal(t, testNode, result)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}

func TestDescribeError(t *testing.T) {
	ctx := context.Background()

	mockRepo := mocks.NewMockNode_repository(t)
	s := &Service{repo: mockRepo}

	testNode := &node.Node{
		Identifier: node.Identifier("test-id"),
	}

	mockRepo.EXPECT().Get(mock.Anything, testNode.Identifier).Return(nil, repository.ErrNodeNotFound)

	result, err := s.Describe(ctx, testNode.Identifier)

	assert.Nil(t, result)
	assert.Error(t, err)

	mockRepo.AssertCalled(t, "Get", ctx, testNode.Identifier)
}
