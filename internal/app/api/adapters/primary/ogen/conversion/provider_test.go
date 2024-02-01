package conversion

import (
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/pkg/gen/texit"
	"github.com/stretchr/testify/assert"
)

func TestTranslateProviderType(t *testing.T) {
	assert.Equal(t, texit.ProviderTypeAWSEcs, TranslateProviderType(provider.TypeAwsEcs))
	assert.Equal(t, texit.ProviderTypeUnknown, TranslateProviderType(provider.TypeUnknown))
}

func TestProviderToSummary(t *testing.T) {
	p := &provider.Provider{
		Name:     provider.Identifier("test-name"),
		Platform: provider.TypeAwsEcs,
	}

	summary := ProviderToSummary(p)

	assert.Equal(t, p.Name.String(), summary.Name)
	assert.Equal(t, TranslateProviderType(p.Platform), summary.Type)
}
