package gateway

import (
	"context"

	"github.com/awlsring/texit/internal/pkg/domain/notification"
)

type Notification interface {
	Type() notification.TopicType
	Endpoint() notification.Endpoint
	SendMessage(ctx context.Context, msg string) error
}
