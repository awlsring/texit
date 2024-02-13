package activity

import (
	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/app/api/ports/service"
)

type Service struct {
	nodeRepo    repository.Node
	execRepo    repository.Execution
	providerGws map[string]gateway.Platform
	tailnetGws  map[string]gateway.Tailnet
}

func NewService(tailnetGws map[string]gateway.Tailnet, providerGws map[string]gateway.Platform, nodeRepo repository.Node, execRepo repository.Execution) service.Activity {
	return &Service{
		nodeRepo:    nodeRepo,
		execRepo:    execRepo,
		providerGws: providerGws,
		tailnetGws:  tailnetGws,
	}
}
