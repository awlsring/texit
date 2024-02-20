package sql_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlNodeRepository) List(ctx context.Context) ([]*node.Node, error) {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Listing nodes from sqlite database")
	var dbn []*NodeSqlRecord
	query := "SELECT * FROM nodes"
	log.Debug().Msgf("Query: %s", query)
	err := r.db.SelectContext(ctx, &dbn, query)
	if err != nil {
		log.Error().Err(err).Msg("Failed to list nodes from sqlite database")
		return nil, err
	}

	log.Debug().Msg("Converting nodes from sqlite")
	nodes := make([]*node.Node, len(dbn))
	for i, n := range dbn {
		nodes[i] = n.ToNode()
	}

	log.Debug().Msgf("Returning nodes: %s", query)
	return nodes, nil
}
