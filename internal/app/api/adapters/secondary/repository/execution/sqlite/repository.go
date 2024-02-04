package sqlite_execution_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	_ "modernc.org/sqlite"
)

type SqliteExecutionRepository struct {
	db interfaces.SqlDatabase
}

func New(db interfaces.SqlDatabase) repository.Execution {
	return &SqliteExecutionRepository{
		db: db,
	}
}

func (r *SqliteExecutionRepository) Init(ctx context.Context) error {
	err := r.initTables(ctx)
	if err != nil {
		return err
	}
	return nil
}
