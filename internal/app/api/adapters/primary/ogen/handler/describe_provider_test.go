package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDescribeProvider(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeProviderParams{
		Name: "test-provider",
	}

	mockProviderSvc := mocks.NewMockProvider_service(t)
	h := New(nil, nil, mockProviderSvc, nil)

	providerId, _ := provider.IdentifierFromString(req.Name)
	testProvider := &provider.Provider{
		Name: providerId,
	}

	mockProviderSvc.EXPECT().Describe(mock.Anything, providerId).Return(testProvider, nil)

	res, err := h.DescribeProvider(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, testProvider.Name.String(), res.(*texit.DescribeProviderResponseContent).Summary.Name)

}

func TestDescribeProviderBadParse(t *testing.T) {
	ctx := context.Background()
	mockProviderSvc := mocks.NewMockProvider_service(t)
	h := New(nil, nil, mockProviderSvc, nil)

	badReq := texit.DescribeProviderParams{
		Name: "",
	}

	res, err := h.DescribeProvider(ctx, badReq)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestDescribeProviderError(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeProviderParams{
		Name: "test-provider",
	}

	mockProviderSvc := mocks.NewMockProvider_service(t)
	h := New(nil, nil, mockProviderSvc, nil)

	providerId, _ := provider.IdentifierFromString(req.Name)

	mockProviderSvc.EXPECT().Describe(ctx, providerId).Return(nil, errors.New("test error"))

	res, err := h.DescribeProvider(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
