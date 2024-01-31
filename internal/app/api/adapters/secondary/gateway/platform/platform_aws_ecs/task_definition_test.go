package platform_aws_ecs

import (
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/stretchr/testify/assert"
)

func TestExtraArgsTailscal(t *testing.T) {
	tn := &tailnet.Tailnet{
		Name: "tailnet@toes",
		Type: tailnet.TypeTailscale,
	}

	predicatedString := "--advertise-exit-node"
	args := makeExtraArgs(tn)

	assert.Equal(t, predicatedString, *args.Value)
}

func TestExtraArgsHeadscale(t *testing.T) {
	tn := &tailnet.Tailnet{
		Name: "https://headscale.toes.com",
		Type: tailnet.TypeHeadscale,
	}

	predicatedString := "--advertise-exit-node --login-server=https://headscale.toes.com"
	args := makeExtraArgs(tn)

	assert.Equal(t, predicatedString, *args.Value)
}
