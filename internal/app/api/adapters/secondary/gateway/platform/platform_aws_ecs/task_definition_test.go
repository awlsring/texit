package platform_aws_ecs

import (
	"testing"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/stretchr/testify/assert"
)

func TestExtraArgs(t *testing.T) {
	tcs := tailnet.ControlServer("https://tailscale.toes.com")

	predicatedString := "--advertise-exit-node --login-server=https://tailscale.toes.com"
	args := makeExtraArgs(tcs)

	assert.Equal(t, predicatedString, *args.Value)
}
