package sqlite_node_repository

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (r *SqliteNodeRepository) Create(ctx context.Context, node *node.Node) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Creating node in sqlite database")

	query := "INSERT INTO nodes (identifier, platform_identifier, provider_identifier, tailnet_identifier, location, preauth_key) VALUES ($1, $2, $3, $4, $5, $6)"
	log.Debug().Msgf("Query: %s", query)
	_, err := r.db.ExecContext(ctx, query, node.Identifier.String(), node.PlatformIdentifier.String(), node.ProviderIdentifier.String(), node.TailnetIdentifier.String(), node.Location.String(), node.PreauthKey.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to create node in sqlite database")
		return err
	}

	log.Debug().Msg("Node created in sqlite database")
	return nil
}
