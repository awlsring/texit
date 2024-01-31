package apiv1

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
)

func TranslateProviderType(p v1.Provider) provider.Type {
	switch p {
	case v1.Provider_PROVIDER_AWS_ECS:
		return provider.TypeAwsEcs
	default:
		return provider.TypeUnknown
	}
}

func SummaryToProvider(sum *v1.ProviderSummary) (*provider.Provider, error) {
	name, err := provider.IdentifierFromString(sum.Id)
	if err != nil {
		return nil, err
	}

	return &provider.Provider{
		Name:     name,
		Platform: TranslateProviderType(sum.Type),
	}, nil
}
