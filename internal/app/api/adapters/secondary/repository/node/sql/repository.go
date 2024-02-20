package sql_node_repository

import (
	"context"

	"github.com/awlsring/texit/internal/app/api/ports/repository"
	"github.com/awlsring/texit/internal/pkg/interfaces"
	_ "modernc.org/sqlite"
)

type SqlNodeRepository struct {
	db interfaces.SqlDatabase
}

func New(db interfaces.SqlDatabase) repository.Node {
	return &SqlNodeRepository{
		db: db,
	}
}

func (r *SqlNodeRepository) Close() {
	r.db.Close()
}

func (r *SqlNodeRepository) Init(ctx context.Context) error {
	err := r.initTables(ctx)
	if err != nil {
		return err
	}
	return nil
}
