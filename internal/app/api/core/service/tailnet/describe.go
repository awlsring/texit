package tailnet

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (s *Service) Describe(ctx context.Context, name tailnet.Identifier) (*tailnet.Tailnet, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msgf("describing tailnet %s", name.String())
	p, ok := s.tailnetMap[name]
	if !ok {
		log.Warn().Msg("unknown tailnet")
		return nil, errors.Wrap(service.ErrUnknownProvider, name.String())
	}

	log.Debug().Msg("returning tailnet")
	return p, nil
}
