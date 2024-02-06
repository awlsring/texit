package node

import (
	"context"
	"sync"

	"github.com/awlsring/texit/internal/app/ui/core/domain/node"
	cnode "github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (s *Service) ListNodes(ctx context.Context) ([]*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("listing nodes")

	nodes, err := s.apiGw.ListNodes(ctx)
	if err != nil {
		log.Error().Err(err).Msg("failed to list nodes")
		return nil, err
	}

	var nodeChan = make(chan *node.Node, len(nodes))
	var wg sync.WaitGroup
	wg.Add(len(nodes))
	log.Debug().Msg("concurrently detailing nodes")
	for _, n := range nodes {
		go func(n *cnode.Node) {
			defer wg.Done()
			no := node.NewBaseNode(n)
			if err := s.detailNode(ctx, no); err != nil {
				log.Warn().Err(err).Msg("failed to detail node")
			}
			nodeChan <- no
		}(n)
	}

	go func() {
		wg.Wait()
		close(nodeChan)
	}()

	ns := make([]*node.Node, 0, len(nodes))
	log.Debug().Msg("collecting detailed nodes")
	for n := range nodeChan {
		ns = append(ns, n)
	}

	log.Debug().Msg("done listing nodes")
	return ns, nil
}
