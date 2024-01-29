package tailscale_gateway

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/tailscale/tailscale-client-go/tailscale"
)

func setAutoApprover(acl *tailscale.ACL) bool {
	exitNodes := acl.AutoApprovers.ExitNode
	for _, exitNode := range exitNodes {
		if exitNode == tagCloudExitNode {
			return false
		}
	}
	acl.AutoApprovers.ExitNode = append(exitNodes, tagCloudExitNode)
	return true
}

func setTagOwners(acl *tailscale.ACL, user string) bool {
	owners := acl.TagOwners[tagCloudExitNode]
	for _, owner := range owners {
		if owner == user {
			return false
		}
	}
	acl.TagOwners[tagCloudExitNode] = append(owners, user)
	return true
}

func (g *TailscaleGateway) updateAcl(ctx context.Context) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Checking if ACL needs update")

	log.Debug().Msg("Getting ACL")
	acl, err := g.client.ACL(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get ACL")
		return err
	}

	log.Debug().Msg("Setting tag owner in ACL")
	ownersUpdated := setTagOwners(acl, g.user)

	log.Debug().Msg("Configuring tag as auto approver")
	approversUpdated := setAutoApprover(acl)

	if !ownersUpdated && !approversUpdated {
		log.Debug().Msg("ACL does not need update")
		return nil
	}

	log.Debug().Msg("Validating ACL")
	err = g.client.ValidateACL(ctx, *acl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to validate ACL")
		return err
	}

	log.Debug().Msg("Updating ACL")
	err = g.client.SetACL(ctx, *acl)
	if err != nil {
		log.Error().Err(err).Msg("Failed to update ACL")
		return err
	}

	log.Debug().Msg("ACL updated")
	return nil
}
