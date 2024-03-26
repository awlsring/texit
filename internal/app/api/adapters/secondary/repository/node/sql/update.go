package sql_node_repository

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlNodeRepository) Update(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Updating node in sqlite database")

	query := "UPDATE nodes SET platform_identifier = $2, provider_identifier = $3, tailnet_identifier = $4, tailnet_device_name = $5, tailnet = $6, location = $7, preauth_key = $8, ephemeral = $9, size = $10, updated_at = $11, provisioning_status = $12 WHERE identifier = $1"
	log.Debug().Msgf("Query: %s", query)
	_, err := r.db.ExecContext(ctx, query, node.Identifier.String(), node.PlatformIdentifier.String(), node.Provider.String(), node.TailnetIdentifier.String(), node.TailnetName.String(), node.Tailnet.String(), node.Location.String(), node.PreauthKey.String(), node.Ephemeral, node.Size.String(), time.Now(), node.ProvisionStatus.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to update node in sqlite database")
		return err
	}

	log.Debug().Msg("Node updated in sqlite database")
	return nil
}
