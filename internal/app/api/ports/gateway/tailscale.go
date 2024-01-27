package gateway

import "context"

type Tailnet interface {
	CreatePreauthKey(context.Context) (string, error)
	DeletePreauthKey(context.Context, string) error
	DeleteDevice(context.Context, string) error
}
