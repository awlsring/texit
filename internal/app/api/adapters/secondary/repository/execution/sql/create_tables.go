package sql_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlExecutionRepository) initTables(ctx context.Context) error {
	log := logger.FromContext(ctx)
	nodeTable := `
		CREATE TABLE IF NOT EXISTS executions (
			identifier TEXT PRIMARY KEY,
			workflow TEXT NOT NULL,
			status TEXT NOT NULL,
			created_at TIMESTAMP NOT NULL,
			updated_at TIMESTAMP NOT NULL,
			finished_at TIMESTAMP,
			results TEXT
		);`

	log.Debug().Msg("creating execution table")
	_, err := r.db.ExecContext(ctx, nodeTable)
	if err != nil {
		log.Error().Err(err).Msg("failed to create execution table")
		return err
	}

	log.Debug().Msg("execution table created")
	return nil
}
