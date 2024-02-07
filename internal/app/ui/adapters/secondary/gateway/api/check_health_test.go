package api_gateway

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestHealthCheck(t *testing.T) {

	t.Run("returns nil when health check is successful", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		apiGateway := New(mockClient)

		mockClient.EXPECT().Health(mock.Anything).Return(&texit.HealthResponseContent{}, nil)

		err := apiGateway.HealthCheck(context.Background())

		assert.NoError(t, err)
		mockClient.AssertExpectations(t)
	})

	t.Run("returns error when health check fails", func(t *testing.T) {
		mockClient := mocks.NewMockInvoker_texit(t)
		apiGateway := New(mockClient)

		mockClient.On("Health", mock.Anything).Return(nil, errors.New("health check failed"))

		err := apiGateway.HealthCheck(context.Background())

		assert.Error(t, err)
		assert.ErrorIs(t, err, gateway.ErrInternalServerError)
		mockClient.AssertExpectations(t)
	})
}
