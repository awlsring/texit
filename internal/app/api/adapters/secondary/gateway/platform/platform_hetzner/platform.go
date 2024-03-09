package platform_hetzner

import (
	"strconv"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

type PlatformHetzner struct {
	client *hcloud.Client
}

func convertPlatformId(p node.PlatformIdentifier) (int64, error) {
	id, err := strconv.ParseInt(p.String(), 10, 64)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(client *hcloud.Client) gateway.Platform {
	return &PlatformHetzner{
		client: client,
	}
}
