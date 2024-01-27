package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
)

func (g *PlatformAwsEcsGateway) GetStatus(ctx context.Context, node *node.Node) (node.Status, error) {
	panic("implement me")
}
