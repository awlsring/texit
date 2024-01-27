package conversion

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func ProviderToSummary(provider *provider.Provider) *teen.ProviderSummary {
	return &teen.ProviderSummary{
		Id:       provider.Name.String(),
		Platform: provider.Platform.String(),
		Default:  provider.Default,
	}
}
