package handler

import (
	"context"
	"fmt"

	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/adapters/primary/cli/flag"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/node"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/pkg/domain/provider"
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
		fmt.Printf("Node - Id: %s | Provider: %s | Location: %s", n.Identifier.String(), n.ProviderIdentifier.String(), n.Location.String())
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

	fmt.Printf("Node - Id: %s | Provider: %s | Location: %s", n.Identifier.String(), n.ProviderIdentifier.String(), n.Location.String())
	return nil
}

func (h *Handler) ProvisionNode(c *cli.Context) error {
	prov, err := provider.IdentifierFromString(c.String(flag.ProviderName))
	if err != nil {
		return err
	}

	location := provider.Location(c.String(flag.ProviderLocation))
	if err != nil {
		return err
	}

	exId, err := h.apiSvc.ProvisionNode(context.Background(), prov, location)
	if err != nil {
		return err
	}

	fmt.Printf("Sent provision node request. Execution Id: %s", exId.String())
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

	fmt.Printf("Sent deprovision node request. Execution Id: %s", exId.String())
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

	fmt.Printf("Sent start node request for node %s", node.String())
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

	fmt.Printf("Sent stop node request for node %s", node.String())
	return nil
}
