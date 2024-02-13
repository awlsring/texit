package step_functions_workflow

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func (s *StepFunctionsWorkflow) scheduleWorkflow(ctx context.Context, wrkArn string, input workflow.ExecutionInput) error {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Marshalling input to JSON")
	inJson, err := input.ToJson()
	if err != nil {
		log.Error().Err(err).Msg("Failed to create workflow input")
		return err
	}

	log.Debug().Msg("Scheduling execution with step fuctions")
	_, err = s.client.StartExecution(ctx, &sfn.StartExecutionInput{
		Name:            aws.String(input.ExecutionIdentifier()),
		StateMachineArn: &wrkArn,
		Input:           aws.String(inJson),
	})
	if err != nil {
		log.Error().Err(err).Msg("Failed to start execution")
		return err
	}

	return nil
}
