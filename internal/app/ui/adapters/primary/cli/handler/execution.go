package handler

import (
	"fmt"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/workflow"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func (h *Handler) DescribeExecution(c *cli.Context) error {
	id, err := workflow.ExecutionIdentifierFromString(c.String(flag.ExecutionId))
	if err != nil {
		e := errors.Wrap(err, "failed to parse execution id")
		return e
	}

	exec, err := h.apiSvc.GetExecution(c.Context, id)
	if err != nil {
		e := errors.Wrap(err, "failed to describe execution")
		return e
	}

	fmt.Printf("Execution - Id: %s | Workflow: %s | Status: %s\n", exec.Identifier.String(), exec.Workflow.String(), exec.Status.String())
	return nil
}
