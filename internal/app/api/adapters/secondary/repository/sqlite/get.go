package sqlite_node_repository

import (
	"context"
	"database/sql"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/api/ports/repository"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (r *SqliteNodeRepository) Get(ctx context.Context, id node.Identifier) (*node.Node, error) {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Getting node from sqlite")
	var ndb NodeSqlRecord
	query := "SELECT * FROM nodes WHERE identifier = $1"
	log.Debug().Msgf("Query: %s", query)
	err := r.db.GetContext(ctx, &ndb, query, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Msgf("Node not found %s", id.String())
			return nil, errors.Wrap(repository.ErrNodeNotFound, id.String())
		}
		log.Error().Err(err).Msg("Failed to get node from sqlite")
		return nil, err
	}

	log.Debug().Msg("Converting node from sqlite")
	n := ndb.ToNode()

	log.Debug().Msgf("Returning node: %s", n.Identifier.String())
	return n, nil
}
