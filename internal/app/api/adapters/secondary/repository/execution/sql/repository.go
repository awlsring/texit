package sql_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/interfaces"
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
	return nil
}

func (r *SqlExecutionRepository) Close() {
	r.db.Close()
}
