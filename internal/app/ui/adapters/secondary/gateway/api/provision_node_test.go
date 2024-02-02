package api_gateway

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

func TestApiGateway_ProvisionNode(t *testing.T) {
	ctx := context.Background()

	provId, _ := provider.IdentifierFromString("test-provider")
	loc := provider.Location("us-east-1")
	tnId, _ := tailnet.IdentifierFromString("test-tailnet")
	eph := true

	testExecutionId := workflow.ExecutionIdentifier("test-execution")

	t.Run("returns execution id when provision node is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ProvisionNode(ctx, &texit.ProvisionNodeRequestContent{
			Provider:  provId.String(),
			Location:  loc.String(),
			Tailnet:   tnId.String(),
			Ephemeral: texit.OptBool{},
		}).Return(&texit.ProvisionNodeResponseContent{
			Execution: testExecutionId.String(),
		}, nil)

		executionId, err := g.ProvisionNode(ctx, provId, loc, tnId, eph)

		assert.NoError(t, err)
		assert.Equal(t, testExecutionId, executionId)
	})

	t.Run("returns error when provision node fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ProvisionNode(ctx, &texit.ProvisionNodeRequestContent{
			Provider:  provId.String(),
			Location:  loc.String(),
			Tailnet:   tnId.String(),
			Ephemeral: texit.OptBool{},
		}).Return(nil, errors.New("provision node failed"))

		executionId, err := g.ProvisionNode(ctx, provId, loc, tnId, eph)

		assert.Error(t, err)
		assert.EqualError(t, err, "provision node failed")
		assert.Equal(t, workflow.ExecutionIdentifier(""), executionId)
	})

	t.Run("returns error when execution id conversion fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		g := ApiGateway{client: mockClient}
		mockClient.EXPECT().ProvisionNode(ctx, &texit.ProvisionNodeRequestContent{
			Provider:  provId.String(),
			Location:  loc.String(),
			Tailnet:   tnId.String(),
			Ephemeral: texit.OptBool{},
		}).Return(&texit.ProvisionNodeResponseContent{
			Execution: "",
		}, nil)

		executionId, err := g.ProvisionNode(ctx, provId, loc, tnId, eph)

		assert.Error(t, err)
		assert.Equal(t, workflow.ExecutionIdentifier(""), executionId)
	})
}
