package platform_aws_ecs

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (g *PlatformAwsEcsGateway) DescribeNode(ctx context.Context, node *node.Node) (*node.Node, error) {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Describing node")
	panic("implement me")
}
