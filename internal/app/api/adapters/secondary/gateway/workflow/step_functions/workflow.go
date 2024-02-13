package step_functions_workflow

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

type StepFunctionsWorkflow struct {
	provisionNodeWorkflowArn   string
	deprovisionNodeWorkflowArn string
	client                     *sfn.Client
}

func New(p, d string, client *sfn.Client) gateway.Workflow {
	return &StepFunctionsWorkflow{
		provisionNodeWorkflowArn:   p,
		deprovisionNodeWorkflowArn: d,
		client:                     client,
	}
}
