package sql_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlNodeRepository) initTables(ctx context.Context) error {
	log := logger.FromContext(ctx)
	nodeTable := `
		CREATE TABLE IF NOT EXISTS nodes (
			identifier TEXT PRIMARY KEY,
			platform_identifier TEXT NOT NULL,
			provider_identifier TEXT NOT NULL,
			tailnet_identifier TEXT NOT NULL,
			tailnet_device_name TEXT NOT NULL,
			tailnet TEXT NOT NULL,
			location TEXT NOT NULL,
			preauth_key TEXT NOT NULL,
			size TEXT NOT NULL,
			ephemeral BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL
		);`

	log.Debug().Msg("creating nodes table")
	_, err := r.db.ExecContext(ctx, nodeTable)
	if err != nil {
		log.Error().Err(err).Msg("failed to create nodes table")
		return err
	}

	log.Debug().Msg("nodes table created")
	return nil
}
