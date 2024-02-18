package pending_execution

import (
	"context"

	tempest "github.com/Amatsagu/Tempest"
	"github.com/awlsring/texit/internal/pkg/logger"
)

type InMemoryTracker struct {
	exs map[string]tempest.Snowflake
}

func NewInMemoryTracker() Tracker {
	return &InMemoryTracker{
		exs: make(map[string]tempest.Snowflake),
	}
}

func (t *InMemoryTracker) AddExecution(ctx context.Context, id string, user tempest.Snowflake) error {
	log := logger.FromContext(ctx)
	log.Debug().Str("id", id).Str("user", user.String()).Msg("Adding execution")
	t.exs[id] = user
	return nil
}

func (t *InMemoryTracker) RemoveExecution(ctx context.Context, id string) error {
	log := logger.FromContext(ctx)
	log.Debug().Str("id", id).Msg("Removing execution")
	delete(t.exs, id)
	return nil
}

func (t *InMemoryTracker) GetExecution(ctx context.Context, id string) (tempest.Snowflake, error) {
	log := logger.FromContext(ctx)
	log.Debug().Str("id", id).Msg("Getting execution")
	user, ok := t.exs[id]
	if !ok {
		log.Warn().Str("id", id).Msg("Execution not found")
		return 0, ErrExecutionNotFound
	}
	return user, nil
}
