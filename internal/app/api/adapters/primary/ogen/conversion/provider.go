package conversion

import (
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func TranslateProviderType(t provider.Type) texit.ProviderType {
	switch t {
	case provider.TypeAwsEcs:
		return texit.ProviderTypeAWSEcs
	default:
		return texit.ProviderTypeUnknown
	}
}

func ProviderToSummary(p *provider.Provider) texit.ProviderSummary {
	return texit.ProviderSummary{
		Name: p.Name.String(),
		Type: TranslateProviderType(p.Platform),
	}
}
