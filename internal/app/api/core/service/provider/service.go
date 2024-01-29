package provider

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/service"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
)

type Service struct {
	provMap map[provider.Identifier]*provider.Provider
}

func NewService(providers []*provider.Provider) service.Provider {
	s := &Service{}

	provMap := make(map[provider.Identifier]*provider.Provider)
	for _, p := range providers {
		provMap[p.Name] = p
	}

	s.provMap = provMap
	return s
}
