package tailnet

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/app/ui/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

const (
	defaultExpiration = 5 * time.Minute
)

func (s *Service) refreshTailnets(ctx context.Context) error {
	s.lastRefresh = time.Now()
	tailnets, err := s.apiGw.ListTailnets(ctx)
	if err != nil {
		return err
	}
	for _, t := range tailnets {
		s.tailnets[string(t.Name.String())] = t
	}
	return nil
}

func (s *Service) getTailnetFromMap(ctx context.Context, id tailnet.Identifier) (*tailnet.Tailnet, error) {
	s.mut.Lock()
	if time.Since(s.lastRefresh) > defaultExpiration {
		if err := s.refreshTailnets(ctx); err != nil {
			return nil, err
		}
	}
	s.mut.Unlock()
	p, ok := s.tailnets[string(id.String())]
	if !ok {
		return nil, service.ErrUnknownTailnet
	}
	return p, nil
}
