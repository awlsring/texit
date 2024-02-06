package node

import (
	"context"
	"sync"

	"github.com/awlsring/texit/internal/app/ui/core/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) detailNode(ctx context.Context, n *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("detailing node")

	var wg sync.WaitGroup
	wg.Add(3)

	log.Debug().Msg("describing provider")
	go func() {
		defer wg.Done()
		p, err := s.provSvc.DescribeProvider(ctx, n.Provider)
		if err != nil {
			log.Warn().Err(err).Msg("failed to describe provider")
			return
		}
		n.ProviderType = p.Platform
	}()

	log.Debug().Msg("describing tailnet")
	go func() {
		defer wg.Done()
		t, err := s.tailSvc.DescribeTailnet(ctx, n.Tailnet)
		if err != nil {
			log.Warn().Err(err).Msg("failed to describe tailnet")
			return
		}
		n.TailnetType = t.Type
	}()

	log.Debug().Msg("getting status")
	go func() {
		defer wg.Done()
		status, err := s.apiGw.GetNodeStatus(ctx, n.Identifier)
		if err != nil {
			log.Warn().Err(err).Msg("failed to get node status")
			return
		}
		n.Status = status
	}()

	log.Debug().Msg("waiting for goroutines to finish")
	wg.Wait()
	log.Debug().Msg("done detailing node")
	return nil
}
