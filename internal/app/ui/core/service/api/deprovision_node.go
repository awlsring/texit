package api

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) DeprovisionNode(ctx context.Context, node node.Identifier) (workflow.ExecutionIdentifier, error) {
	log := logger.FromContext(ctx)
	resp, err := s.apiGw.DeprovisionNode(ctx, node)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			log.Warn().Err(err).Msg("node not found")
			return "", service.ErrUnknownNode
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			log.Warn().Err(err).Msg("invalid input")
			return "", service.ErrInvalidInputError
		}
		log.Error().Err(err).Msg("failed to deprovision node")
		return "", err
	}
	return resp, nil
}
