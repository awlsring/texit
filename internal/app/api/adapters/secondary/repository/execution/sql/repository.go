package sql_execution_repository

import (
	"context"
	"time"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	"github.com/rs/zerolog/log"
	_ "modernc.org/sqlite"
)

type SqlExecutionRepository struct {
	db interfaces.SqlDatabase
}

func New(db interfaces.SqlDatabase) repository.Execution {
	return &SqlExecutionRepository{
		db: db,
	}
}

func (r *SqlExecutionRepository) Init(ctx context.Context) error {
	err := r.initTables(ctx)
	if err != nil {
		return err
	}
	err = r.dumpOldExecutions(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (r *SqlExecutionRepository) dumpOldExecutions(ctx context.Context) error {
	daysAgo5 := time.Now().AddDate(0, 0, -5)
	query := "DELETE FROM executions WHERE finished_at < $1"
	_, err := r.db.ExecContext(ctx, query, daysAgo5)
	if err != nil {
		log.Error().Err(err).Msg("Failed to delete node from sqlite database")
		return err
	}
	return nil
}

func (r *SqlExecutionRepository) Close() {
	r.db.Close()
}
