package sql_execution_repository

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlExecutionRepository) CloseExecution(ctx context.Context, id workflow.ExecutionIdentifier, result workflow.Status, output workflow.SerializedExecutionResult) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("Closing execution in sqlite")

	query := "UPDATE executions SET status = $1, updated_at = $2, finished_at = $3, results = $4 WHERE identifier = $5"
	log.Debug().Msgf("Query: %s", query)
	now := time.Now()
	_, err := r.db.ExecContext(ctx, query, result.String(), now, now, output.String(), id.String())
	if err != nil {
		log.Error().Err(err).Msg("Failed to close execution in sqlite database")
		return err
	}

	log.Debug().Msg("Execution closed in sqlite database")
	return nil
}