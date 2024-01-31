package provider

import (
	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
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
