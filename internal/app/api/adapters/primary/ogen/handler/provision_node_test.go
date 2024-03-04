package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestProvisionNode(t *testing.T) {
	ctx := context.Background()
	req := &texit.ProvisionNodeRequestContent{
		Provider:  "test-provider",
		Tailnet:   "test-tailnet",
		Location:  "us-east-1",
		Ephemeral: texit.OptBool{Value: true},
	}

	mockProviderSvc := mocks.NewMockProvider_service(t)
	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, mockProviderSvc, mockTailnetSvc, nil, nil)

	provId, _ := provider.IdentifierFromString(req.Provider)
	testProvider := &provider.Provider{
		Name:     provId,
		Platform: provider.TypeAwsEcs,
	}

	tnId, _ := tailnet.IdentifierFromString(req.Tailnet)
	testTailnet := &tailnet.Tailnet{
		Name: tnId,
	}

	testLocation := provider.Location("us-east-1")

	testExecutionId := workflow.ExecutionIdentifier("test-execution")

	mockProviderSvc.EXPECT().Describe(ctx, provId).Return(testProvider, nil)
	mockTailnetSvc.EXPECT().Describe(ctx, tnId).Return(testTailnet, nil)
	mockWorkSvc.EXPECT().LaunchProvisionNodeWorkflow(ctx, testProvider, testLocation, testTailnet, req.Ephemeral.Value).Return(testExecutionId, nil)

	res, err := h.ProvisionNode(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, testExecutionId.String(), res.(*texit.ProvisionNodeResponseContent).Execution)
}

func TestProvisionNodeError(t *testing.T) {
	ctx := context.Background()
	req := &texit.ProvisionNodeRequestContent{
		Provider:  "test-provider",
		Tailnet:   "test-tailnet",
		Location:  "us-east-1",
		Ephemeral: texit.OptBool{Value: true},
	}

	mockProviderSvc := mocks.NewMockProvider_service(t)
	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, mockProviderSvc, mockTailnetSvc, nil, nil)

	provId, _ := provider.IdentifierFromString(req.Provider)

	mockProviderSvc.EXPECT().Describe(ctx, provId).Return(nil, errors.New("test error"))

	res, err := h.ProvisionNode(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
