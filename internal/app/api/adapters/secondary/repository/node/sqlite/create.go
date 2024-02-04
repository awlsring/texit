package sqlite_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqliteNodeRepository) Create(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node in sqlite database")

	query := "INSERT INTO nodes (identifier, platform_identifier, provider_identifier, tailnet_identifier, tailnet_device_name, tailnet, location, preauth_key, ephemeral, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	log.Debug().Msgf("Query: %s", query)
	_, err := r.db.ExecContext(ctx, query, node.Identifier.String(), node.PlatformIdentifier.String(), node.Provider.String(), node.TailnetIdentifier.String(), node.TailnetName.String(), node.Tailnet.String(), node.Location.String(), node.PreauthKey.String(), node.Ephemeral, node.CreatedAt, node.UpdatedAt)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node in sqlite database")
		return err
	}

	log.Debug().Msg("Node created in sqlite database")
	return nil
}
