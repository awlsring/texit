package interfaces

import (
	"context"

	"github.com/linode/linodego"
)

type LinodeClient interface {
	CreateStackscript(ctx context.Context, opts linodego.StackscriptCreateOptions) (*linodego.Stackscript, error)
	ListStackscripts(ctx context.Context, opts *linodego.ListOptions) ([]linodego.Stackscript, error)
	DeleteStackscript(ctx context.Context, scriptID int) error
	CreateInstance(ctx context.Context, opts linodego.InstanceCreateOptions) (*linodego.Instance, error)
	DeleteInstance(ctx context.Context, linodeID int) error
	GetInstance(ctx context.Context, linodeID int) (*linodego.Instance, error)
	BootInstance(ctx context.Context, linodeID int, configID int) error
	ShutdownInstance(ctx context.Context, id int) error
}
