package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/pkg/errors"
)

func (s *Service) ProvisionNode(ctx context.Context, prov provider.Identifier, loc provider.Location, tn tailnet.Identifier, sz node.Size, e bool) (workflow.ExecutionIdentifier, error) {
	resp, err := s.apiGw.ProvisionNode(ctx, prov, loc, tn, sz, e)
	if err != nil {
		if errors.Is(err, gateway.ErrInvalidInputError) {
			return "", service.ErrInvalidInputError

		}
		return "", err
	}

	return resp, nil
}
