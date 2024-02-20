package sql_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
)

func (r *SqlExecutionRepository) CreateExecution(ctx context.Context, ex *workflow.Execution) error {
	log := logger.FromContext(ctx)
	log.Debug().Msg("creating execution in repo")

	query := "INSERT INTO executions (identifier, workflow, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)"
	log.Debug().Msgf("Query: %s", query)
	_, err := r.db.ExecContext(ctx, query, ex.Identifier.String(), ex.Workflow.String(), ex.Status.String(), ex.Created, ex.Updated)
	if err != nil {
		log.Error().Err(err).Msg("Failed to create execution in sqlite database")
		return err
	}

	log.Debug().Msg("Execution created in sqlite database")
	return nil
}
