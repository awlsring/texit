package sqlite_node_repository

import (
	"context"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/core/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
)

func (r *SqliteNodeRepository) Delete(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting node from sqlite database")

	query := "DELETE FROM nodes WHERE id = $1"
	log.Debug().Msgf("Query: %s", query)
	_, err := r.db.ExecContext(ctx, query, id.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node from sqlite database")
		return err
	}

	log.Debug().Msg("Node deleted from sqlite database")
	return nil
}
