package handler

import (
	"context"
	"fmt"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/tailnet"
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
		fmt.Printf("Node - Id: %s | Provider: %s | Location: %s\n", n.Identifier.String(), n.Provider.String(), n.Location.String())
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

	fmt.Printf("Node - Id: %s | Provider: %s | Location: %s\n", n.Identifier.String(), n.Provider.String(), n.Location.String())
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
