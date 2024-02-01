package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/stretchr/testify/assert"
)

func TestListTailnets(t *testing.T) {
	ctx := context.Background()

	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	h := New(nil, nil, nil, mockTailnetSvc)

	testTailnets := []*tailnet.Tailnet{
		{
			Name: tailnet.Identifier("test-tailnet-1"),
		},
		{
			Name: tailnet.Identifier("test-tailnet-2"),
		},
	}

	mockTailnetSvc.EXPECT().List(ctx).Return(testTailnets, nil)

	res, err := h.ListTailnets(ctx)

	assert.NoError(t, err)
	assert.Len(t, res.Summaries, len(testTailnets))
}

func TestListTailnetsError(t *testing.T) {
	ctx := context.Background()

	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	h := New(nil, nil, nil, mockTailnetSvc)

	mockTailnetSvc.EXPECT().List(ctx).Return(nil, errors.New("test error"))

	res, err := h.ListTailnets(ctx)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestListTailnetsEmpty(t *testing.T) {
	ctx := context.Background()

	mockTailnetSvc := mocks.NewMockTailnet_service(t)
	h := New(nil, nil, nil, mockTailnetSvc)

	mockTailnetSvc.EXPECT().List(ctx).Return([]*tailnet.Tailnet{}, nil)

	res, err := h.ListTailnets(ctx)

	assert.NoError(t, err)
	assert.Len(t, res.Summaries, 0)
}
