package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestDescribeTailnet(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeTailnetParams{
		Name: "test-tailnet",
	}

	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	h := New(nil, nil, nil, mockTailnetSvc, nil)

	tailnetId, _ := tailnet.IdentifierFromString(req.Name)
	testTailnet := &tailnet.Tailnet{
		Name: tailnetId,
	}

	mockTailnetSvc.EXPECT().Describe(mock.Anything, tailnetId).Return(testTailnet, nil)

	res, err := h.DescribeTailnet(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, testTailnet.Name.String(), res.(*texit.DescribeTailnetResponseContent).Summary.Name)
}

func TestDescribeTailnetFailToParse(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeTailnetParams{
		Name: "test-tailnet",
	}

	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	h := New(nil, nil, nil, mockTailnetSvc, nil)

	tailnetId, _ := tailnet.IdentifierFromString(req.Name)
	badReq := texit.DescribeTailnetParams{
		Name: "",
	}

	res, err := h.DescribeTailnet(ctx, badReq)

	assert.Error(t, err)
	assert.Nil(t, res)

	t.Run("failed to describe tailnet", func(t *testing.T) {
		mockTailnetSvc.EXPECT().Describe(ctx, tailnetId).Return(nil, errors.New("test error"))

		res, err := h.DescribeTailnet(ctx, req)

		assert.Error(t, err)
		assert.Nil(t, res)
	})
}

func TestDescribeTailnetError(t *testing.T) {
	ctx := context.Background()
	req := texit.DescribeTailnetParams{
		Name: "test-tailnet",
	}

	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	h := New(nil, nil, nil, mockTailnetSvc, nil)

	tailnetId, _ := tailnet.IdentifierFromString(req.Name)

	mockTailnetSvc.EXPECT().Describe(ctx, tailnetId).Return(nil, errors.New("test error"))

	res, err := h.DescribeTailnet(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
