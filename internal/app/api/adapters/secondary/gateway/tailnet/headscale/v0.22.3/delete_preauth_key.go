package headscale_v0_22_3_gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
)

func (g *HeadscaleGateway) DeletePreauthKey(ctx context.Context, key tailnet.PreauthKey) error {
	//TODO uneeded??
	return nil
}
