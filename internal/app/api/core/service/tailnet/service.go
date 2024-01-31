package tailnet

import (
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

type Service struct {
	tailnetMap map[tailnet.Identifier]*tailnet.Tailnet
}

func NewService(tailnets []*tailnet.Tailnet) service.Tailnet {
	s := &Service{}

	tailnetMap := make(map[tailnet.Identifier]*tailnet.Tailnet)
	for _, t := range tailnets {
		tailnetMap[t.Name] = t
	}

	s.tailnetMap = tailnetMap
	return s
}
