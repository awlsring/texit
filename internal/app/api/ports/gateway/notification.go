package gateway

import (
	"context"
)

type Notification interface {
	Endpoint() string
	SendMessage(ctx context.Context, msg string) error
}
