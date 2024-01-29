package api

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
)

func (s *Service) ProvisionNode(ctx context.Context, prov provider.Identifier, loc provider.Location, tn tailnet.Identifier) (workflow.ExecutionIdentifier, error) {
	return s.apiGw.ProvisionNode(ctx, prov, loc, tn)
}
