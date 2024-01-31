// FILEPATH: /Users/awlsring/Code/texit/internal/app/api/adapters/primary/grpc/conversion/provider_conversion_test.go

package conversion

import (
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	teen "github.com/awlsring/texit/pkg/gen/client/v1"
	"github.com/stretchr/testify/assert"
)

func TestTranslateProvider(t *testing.T) {
	assert.Equal(t, teen.Provider_PROVIDER_AWS_ECS, TranslateProvider(provider.TypeAwsEcs))
	assert.Equal(t, teen.Provider_PROVIDER_UNKNOWN_UNSPECIFIED, TranslateProvider(provider.TypeUnknown))
}

func TestProviderToSummary(t *testing.T) {
	p := &provider.Provider{
		Name:     provider.Identifier("test-name"),
		Platform: provider.TypeAwsEcs,
	}

	summary := ProviderToSummary(p)

	assert.Equal(t, p.Name.String(), summary.Id)
	assert.Equal(t, TranslateProvider(p.Platform), summary.Type)
}
