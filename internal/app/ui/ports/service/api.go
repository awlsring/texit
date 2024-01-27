package service

import "context"

type Api interface {
	CheckServerHealth(ctx context.Context) error
}
