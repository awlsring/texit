package node

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/core/domain/node"
	"github.com/awlsring/texit/internal/app/ui/ports/gateway"
	"github.com/awlsring/texit/internal/app/ui/ports/service"
	cnode "github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) DescribeNode(ctx context.Context, id cnode.Identifier) (*node.Node, error) {
	log := logger.FromContext(ctx)
	n, err := s.apiGw.DescribeNode(ctx, id)
	if err != nil {
		if errors.Is(err, gateway.ErrResourceNotFoundError) {
			return nil, service.ErrUnknownNode
		}
		if errors.Is(err, gateway.ErrInvalidInputError) {
			base := errors.Unwrap(err)
			return nil, errors.Wrap(service.ErrInvalidInputError, base.Error())
		}
		return nil, err
	}

	no := node.NewBaseNode(n)
	if err = s.detailNode(ctx, no); err != nil {
		log.Warn().Err(err).Msg("failed to detail node")
	}

	return no, nil
}
