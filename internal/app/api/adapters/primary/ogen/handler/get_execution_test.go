package handler

import (
	"context"
	"errors"
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/mocks"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestSuccessfulGetExecution(t *testing.T) {
	ctx := context.Background()
	req := texit.GetExecutionParams{
		Identifier: "test-execution",
	}

	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, nil, nil, nil)

	exId, _ := workflow.ExecutionIdentifierFromString(req.Identifier)
	testExecution := &workflow.Execution{
		Identifier: exId,
		Status:     workflow.StatusComplete,
	}

	mockWorkSvc.EXPECT().GetExecution(ctx, exId).Return(testExecution, nil)

	res, err := h.GetExecution(ctx, req)

	assert.NoError(t, err)
	assert.Equal(t, testExecution.Identifier.String(), res.(*texit.GetExecutionResponseContent).Summary.Identifier)
}

func TestFailedToParseExecutionId(t *testing.T) {
	ctx := context.Background()
	badReq := texit.GetExecutionParams{
		Identifier: "",
	}

	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, nil, nil, nil)

	res, err := h.GetExecution(ctx, badReq)

	assert.Error(t, err)
	assert.Nil(t, res)
}

func TestFailedToGetExecution(t *testing.T) {
	ctx := context.Background()
	req := texit.GetExecutionParams{
		Identifier: "test-execution",
	}

	mockWorkSvc := mocks.NewMockWorkflow_service(t)
	h := New(nil, mockWorkSvc, nil, nil, nil)

	exId, _ := workflow.ExecutionIdentifierFromString(req.Identifier)

	mockWorkSvc.EXPECT().GetExecution(ctx, exId).Return(nil, errors.New("test error"))

	res, err := h.GetExecution(ctx, req)

	assert.Error(t, err)
	assert.Nil(t, res)
}
