package apiv1

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func SummaryToProvider(sum *v1.ProviderSummary) (*provider.Provider, error) {
	name, err := provider.IdentifierFromString(sum.Id)
	if err != nil {
		return nil, err
	}

	t, err := provider.TypeFromString(sum.Platform)
	if err != nil {
		return nil, err
	}

	return &provider.Provider{
		Name:     name,
		Default:  sum.Default,
		Platform: t,
	}, nil
}
