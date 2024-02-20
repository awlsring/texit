package sql_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlNodeRepository) Delete(ctx context.Context, id node.Identifier) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Deleting node from sqlite database")

	query := "DELETE FROM nodes WHERE identifier = $1"
	log.Debug().Msgf("Query: %s", query)
	_, err := r.db.ExecContext(ctx, query, id.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node from sqlite database")
		return err
	}

	log.Debug().Msg("Node deleted from sqlite database")
	return nil
}
