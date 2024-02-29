package notification

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
)

func (s *Service) ListNotifiers(ctx context.Context) []*notification.Notifier {
	return s.notifiers
}
