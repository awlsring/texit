package gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
)

type Workflow interface {
	ProvisionNode(ctx context.Context, input *workflow.ProvisionNodeInput) error
	DeprovisionNode(ctx context.Context, input *workflow.DeprovisionNodeInput) error
}
