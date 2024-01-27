package conversion

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/provider"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func ProviderToSummary(provider *provider.Provider) *teen.ProviderSummary {
	return &teen.ProviderSummary{} // TODO: Implement
}
