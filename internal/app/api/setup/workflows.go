package setup

import (
	local_workflow "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/workflow/local"
	step_functions_workflow "github.com/awlsring/texit/internal/app/api/adapters/secondary/gateway/workflow/step_functions"
	"github.com/awlsring/texit/internal/app/api/config"
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/appinit"
	cconfig "github.com/awlsring/texit/internal/pkg/config"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/aws/aws-sdk-go-v2/service/sfn"
)

func LoadWorkflowEngine(cfg *config.WorkflowConfig) gateway.Workflow {
	switch cfg.Type {
	case config.WorkflowTypeLocal:
		return LoadLocalWorkflowGateway()
	case config.WorkflowTypeSfn:
		return LoadStepFunctionsWorkflowGateway(cfg)
	default:
		panic("unknown workflow type")
	}
}

func LoadLocalWorkflowGateway() gateway.Workflow {
	workChan := make(chan workflow.ExecutionInput)
	return local_workflow.New(workChan)
}

func LoadStepFunctionsWorkflowGateway(cfg *config.WorkflowConfig) gateway.Workflow {
	awsCfg, err := cconfig.LoadAwsConfig(cfg.AccessKey, cfg.SecretKey, cfg.Region)
	appinit.PanicOnErr(err)
	states := sfn.NewFromConfig(awsCfg)
	return step_functions_workflow.New(cfg.ProvisionWorkflowArn, cfg.DeprovisionWorkflowArn, states)
}
