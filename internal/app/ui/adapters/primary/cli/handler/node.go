package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/texit/internal/pkg/domain/node"
	"github.com/awlsring/texit/internal/pkg/domain/provider"
	"github.com/awlsring/texit/internal/pkg/domain/tailnet"
	"github.com/awlsring/texit/internal/pkg/domain/workflow"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func (h *Handler) ListNodes(c *cli.Context) error {
	nodes, err := h.apiSvc.ListNodes(context.Background())
	if err != nil {
		e := errors.Wrap(err, "failed to list nodes")
		return e
	}

	if len(nodes) == 0 {
		fmt.Println("No nodes found")
		return nil
	}

	fmt.Printf("Nodes: %d\n", len(nodes))
	fmt.Println("==========================")
	for _, n := range nodes {
		fmt.Printf("Node - Id: %s | Provider: %s | Location: %s | Tailnet: %s | TailnetName: %s | TailnetId: %s\n", n.Identifier.String(), n.Provider.String(), n.Location.String(), n.Tailnet.String(), n.TailnetName, n.TailnetIdentifier.String())
	}
	return nil
}

func (h *Handler) DescribeNode(c *cli.Context) error {
	id, err := node.IdentifierFromString(c.String(flag.NodeId))
	if err != nil {
		e := errors.Wrap(err, "failed to parse node id")
		return e
	}

	n, err := h.apiSvc.GetNode(context.Background(), id)
	if err != nil {
		e := errors.Wrap(err, "failed to get node")
		return e
	}

	fmt.Printf("Node - Id: %s | Provider: %s | Location: %s | Tailnet: %s | TailnetName: %s | TailnetId: %s\n", n.Identifier.String(), n.Provider.String(), n.Location.String(), n.Tailnet.String(), n.TailnetName.String(), n.TailnetIdentifier.String())
	return nil
}

func (h *Handler) GetNodeStatus(c *cli.Context) error {
	id, err := node.IdentifierFromString(c.String(flag.NodeId))
	if err != nil {
		e := errors.Wrap(err, "failed to parse node id")
		return e
	}

	status, err := h.apiSvc.GetNodeStatus(context.Background(), id)
	if err != nil {
		e := errors.Wrap(err, "failed to get node status")
		return e
	}

	fmt.Printf("Node - Id: %s | Status: %s\n", id.String(), status.String())
	return nil
}

func (h *Handler) ProvisionNode(c *cli.Context) error {
	prov, err := provider.IdentifierFromString(c.String(flag.Provider))
	if err != nil {
		e := errors.Wrap(err, "failed to parse provider name")
		return e
	}

	location := provider.Location(c.String(flag.ProviderLocation))
	ephemeral := c.Bool(flag.EphemeralNode)

	tn, err := tailnet.IdentifierFromString(c.String(flag.Tailnet))
	if err != nil {
		e := errors.Wrap(err, "failed to parse tailnet name")
		return e
	}

	exId, err := h.apiSvc.ProvisionNode(context.Background(), prov, location, tn, ephemeral)
	if err != nil {
		return err
	}

	fmt.Printf("Sent provision node request. Execution Id: %s\n", exId.String())
	if !c.Bool(flag.NoPollExecution) {
		results, err := h.pollWorkflow(exId, 120, 5)
		if err != nil {
			return err
		}
		for _, r := range results {
			fmt.Println(r)
		}
		fmt.Println("Provision node workflow complete.")
	}
	return nil
}

func (h *Handler) DeprovisionNode(c *cli.Context) error {
	node, err := node.IdentifierFromString(c.String(flag.NodeId))
	if err != nil {
		return err
	}

	exId, err := h.apiSvc.DeprovisionNode(context.Background(), node)
	if err != nil {
		return err
	}

	fmt.Printf("Sent deprovision node request. Execution Id: %s\n", exId.String())

	if !c.Bool(flag.NoPollExecution) {
		_, err := h.pollWorkflow(exId, 10, 2)
		if err != nil {
			return err
		}
		fmt.Printf("Deprovision node workflow for %s complete.\n", node.String())
	}

	return nil
}

func (h *Handler) StartNode(c *cli.Context) error {
	node, err := node.IdentifierFromString(c.String(flag.NodeId))
	if err != nil {
		return err
	}

	err = h.apiSvc.StartNode(context.Background(), node)
	if err != nil {
		return err
	}

	fmt.Printf("Sent start node request for node %s\n", node.String())
	return nil
}

func (h *Handler) StopNode(c *cli.Context) error {
	node, err := node.IdentifierFromString(c.String(flag.NodeId))
	if err != nil {
		return err
	}

	err = h.apiSvc.StopNode(context.Background(), node)
	if err != nil {
		return err
	}

	fmt.Printf("Sent stop node request for node %s\n", node.String())
	return nil
}

func (h *Handler) pollWorkflow(ex workflow.ExecutionIdentifier, intervals, wait int) ([]string, error) {
	fmt.Println("Polling workflow until completion...")
	time.Sleep(time.Duration(wait) * time.Second)
	for i := 0; i < intervals; i++ {
		exec, err := h.apiSvc.GetExecution(context.Background(), ex)
		if err != nil {
			return nil, err
		}

		if exec.Status == workflow.StatusFailed {
			fmt.Println("Workflow failed")
			fmt.Println("Errors:")
			for _, e := range exec.Results {
				fmt.Println(e)
			}
			return nil, errors.New("workflow failed")
		}

		if exec.Status == workflow.StatusComplete {
			return exec.Results, nil
		}

		fmt.Printf("Workflow still running... waiting %v seconds\n", wait)
		time.Sleep(time.Duration(wait) * time.Second)
	}

	return nil, errors.New("workflow timed out")
}
