package sfn_activities

import (
	"context"
	"encoding/json"
	"errors"
)

func stringToInput[I any](input string) (*I, error) {
	var in I
	err := json.Unmarshal([]byte(input), &in)
	if err != nil {
		return nil, err
	}
	return &in, nil
}

func (h *SfnActivityHandler) routeActivity(ctx context.Context, activity, input string) (interface{}, error) {
	switch activity {
	case "deleteNode":
		in, err := stringToInput[DeleteNodeInput](input)
		if err != nil {
			return nil, err
		}
		return h.deleteNodeActivity(ctx, in)
	case "deleteNodeRecord":
		in, err := stringToInput[DeleteNodeRecordInput](input)
		if err != nil {
			return nil, err
		}
		return h.deleteNodeRecordActivity(ctx, in)
	case "removeTailnetDevice":
		in, err := stringToInput[RemoveTailnetDeviceInput](input)
		if err != nil {
			return nil, err
		}
		err = h.removeTailnetDeviceActivity(ctx, in)
		return nil, err
	case "formIdentifiers":
		in, err := stringToInput[FormIdentifiersInput](input)
		if err != nil {
			return nil, err
		}
		return h.formIdentifiersActivity(ctx, in)
	case "createNode":
		in, err := stringToInput[CreateNodeInput](input)
		if err != nil {
			return nil, err
		}
		return h.createNodeActivity(ctx, in)
	case "createNodeRecord":
		in, err := stringToInput[CreateNodeRecordInput](input)
		if err != nil {
			return nil, err
		}
		err = h.createNodeRecordActivity(ctx, in)
		return nil, err
	case "updateNodeRecord":
		in, err := stringToInput[UpdateNodeRecordInput](input)
		if err != nil {
			return nil, err
		}
		err = h.updateNodeRecordActivity(ctx, in)
		return nil, err
	case "createPreauthKey":
		in, err := stringToInput[CreatePreauthKeyInput](input)
		if err != nil {
			return nil, err
		}
		return h.createPreAuthKeyActivity(ctx, in)
	case "enableExitNode":
		in, err := stringToInput[EnableExitNodeInput](input)
		if err != nil {
			return nil, err
		}
		err = h.enableExitNodeActivity(ctx, in)
		return nil, err
	case "getDeviceId":
		in, err := stringToInput[GetDeviceIdInput](input)
		if err != nil {
			return nil, err
		}
		return h.getDeviceIdActivity(ctx, in)
	case "closeExecution":
		in, err := stringToInput[CloseExecutionInput](input)
		if err != nil {
			return nil, err
		}
		err = h.closeExecutionActivity(ctx, in)
		return nil, err
	default:
		return nil, errors.New("unknown activity")
	}
}
