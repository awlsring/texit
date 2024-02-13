package sfn_activities

import (
	"context"
	"encoding/json"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/rs/zerolog"
)

type SfnActivityHandler struct {
	actSvc   service.Activity
	logLevel zerolog.Level
}

func New(actSvc service.Activity) *SfnActivityHandler {
	return &SfnActivityHandler{
		actSvc: actSvc,
	}
}

type ActivityRequest struct {
	ActivityName string      `json:"activityName"`
	Input        interface{} `json:"input"`
}

func (h *SfnActivityHandler) HandleRequest(ctx context.Context, input ActivityRequest) (interface{}, error) {
	ctx = logger.InitContextLogger(ctx, h.logLevel)

	inputRaw, err := json.Marshal(input.Input)
	if err != nil {
		return nil, err
	}

	return h.routeActivity(ctx, input.ActivityName, string(inputRaw))
}
