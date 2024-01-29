package conversion

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	teen "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func TranslateProvider(p provider.Type) teen.Provider {
	switch p {
	case provider.TypeAwsEcs:
		return teen.Provider_PROVIDER_AWS_ECS
	default:
		return teen.Provider_PROVIDER_UNKNOWN_UNSPECIFIED
	}
}

func ProviderToSummary(provider *provider.Provider) *teen.ProviderSummary {
	return &teen.ProviderSummary{
		Id:   provider.Name.String(),
		Type: TranslateProvider(provider.Platform),
	}
}
