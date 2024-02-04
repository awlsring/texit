package sqlite_execution_repository

import (
	"context"
	"database/sql"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/awlsring/texit/internal/pkg/logger"
	"github.com/pkg/errors"
)

func (r *SqliteExecutionRepository) GetExecution(ctx context.Context, id workflow.ExecutionIdentifier) (*workflow.Execution, error) {
	log := logger.FromContext(ctx)

	log.Debug().Msg("Getting execution from sqlite")
	var ndb ExecutionSqlRecord
	query := "SELECT * FROM executions WHERE identifier = $1"
	log.Debug().Msgf("Query: %s", query)
	err := r.db.GetContext(ctx, &ndb, query, id.String())
	if err != nil {
		if err == sql.ErrNoRows {
			log.Warn().Msgf("Execution not found %s", id.String())
			return nil, errors.Wrap(repository.ErrNodeNotFound, id.String())
		}
		log.Error().Err(err).Msg("Failed to get execution from sqlite")
		return nil, err
	}

	log.Debug().Msg("Converting execution from record")
	n := ndb.ToExecution()

	log.Debug().Msgf("Returning execution: %s", n.Identifier.String())
	return n, nil
}
