package pending_execution

import (
	"context"
	"errors"

	tempest "github.com/Amatsagu/Tempest"
)

var (
	ErrExecutionNotFound = errors.New("execution not found")
)

type Tracker interface {
	AddExecution(ctx context.Context, id string, user tempest.Snowflake) error
	RemoveExecution(ctx context.Context, id string) error
	GetExecution(ctx context.Context, id string) (tempest.Snowflake, error)
}
