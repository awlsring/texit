package api

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
)

func (s *Service) ListNotifiers(ctx context.Context) ([]*notification.Notifier, error) {
	return s.apiGw.ListNotifiers(ctx)
}
