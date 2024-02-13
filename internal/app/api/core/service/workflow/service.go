package workflow

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/repository"
)

type Service struct {
	nodeRepo repository.Node
	excRepo  repository.Execution
	wfGw     gateway.Workflow
}

func NewService(nodeRepo repository.Node, excRepo repository.Execution, wfGw gateway.Workflow) *Service {
	return &Service{
		nodeRepo: nodeRepo,
		excRepo:  excRepo,
		wfGw:     wfGw,
	}
}
