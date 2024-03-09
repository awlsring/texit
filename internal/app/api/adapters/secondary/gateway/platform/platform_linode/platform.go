package platform_linode

import (
	"strconv"

	"github.com/awlsring/texit/internal/app/api/ports/gateway"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/interfaces"
)

const (
	DefaultSmallInstanceType  = "g6-nanode-1"
	DefaultMediumInstanceType = "g6-standard-1"
	DefaultLargeInstanceType  = "g6-standard-2"
)

type PlatformLinode struct {
	client interfaces.LinodeClient
}

func convertPlatformId(p node.PlatformIdentifier) (int, error) {
	id, err := strconv.Atoi(p.String())
	if err != nil {
		return 0, err
	}
	return id, nil
}

func New(client interfaces.LinodeClient) gateway.Platform {
	return &PlatformLinode{
		client: client,
	}
}
