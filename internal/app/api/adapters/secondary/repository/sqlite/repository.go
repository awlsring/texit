package sqlite_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	_ "github.com/mattn/go-sqlite3"
)

type SqliteNodeRepository struct {
	db interfaces.SqlDatabase
}

func New(db interfaces.SqlDatabase) repository.Node {
	return &SqliteNodeRepository{
		db: db,
	}
}

func (r *SqliteNodeRepository) Init(ctx context.Context) error {
	err := r.initTables(ctx)
	if err != nil {
		return err
	}
	return nil
}
