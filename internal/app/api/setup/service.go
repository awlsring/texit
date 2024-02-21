package setup

import (
	"github.com/awlsring/texit/internal/app/api/config"
	provSvc "github.com/awlsring/texit/internal/app/api/core/service/provider"
	tailnetSvc "github.com/awlsring/texit/internal/app/api/core/service/tailnet"

	"github.com/awlsring/texit/internal/app/api/ports/service"
	"github.com/awlsring/texit/internal/pkg/appinit"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

func LoadProviderService(providers []*config.ProviderConfig) service.Provider {
	provs := []*provider.Provider{}
	for _, p := range providers {
		name, err := provider.IdentifierFromString(p.Name)
		appinit.PanicOnErr(err)
		typ, err := provider.TypeFromString(p.Type.String())
		appinit.PanicOnErr(err)
		provs = append(provs, &provider.Provider{
			Name:     name,
			Platform: typ,
		})
	}
	svc := provSvc.NewService(provs)
	return svc
}

func LoadTailnetService(tailnets []*config.TailnetConfig) service.Tailnet {
	provs := []*tailnet.Tailnet{}
	for _, t := range tailnets {
		name, err := tailnet.IdentifierFromString(t.Tailnet)
		appinit.PanicOnErr(err)
		typ, err := tailnet.TypeFromString(t.Type.String())
		appinit.PanicOnErr(err)
		cs, err := tailnet.ControlServerFromString(t.ControlServer)
		appinit.PanicOnErr(err)
		provs = append(provs, &tailnet.Tailnet{
			Name:          name,
			Type:          typ,
			ControlServer: cs,
		})
	}
	svc := tailnetSvc.NewService(provs)
	return svc
}
