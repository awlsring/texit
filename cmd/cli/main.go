package main

import (
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/adapters/primary/cli"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/adapters/primary/cli/handler"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/config"
	"github.com/awlsring/tailscale-cloud-exit-nodes/internal/app/ui/core/service/api"
	v1 "github.com/awlsring/tailscale-cloud-exit-nodes/pkg/gen/client/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

func initClient(address string) v1.TailscaleEphemeralExitNodesServiceClient {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	panicOnErr(err)
	return v1.NewTailscaleEphemeralExitNodesServiceClient(conn)
}

func main() {
	cfg, err := config.LoadFromFile("cli-config.yaml")
	panicOnErr(err)
	client := initClient(cfg.Api.Address)

	svc := api.NewService(cfg.Api.ApiKey, client)
	hdl := handler.New(svc)
	tool := cli.New(hdl)
	panicOnErr(tool.Run())
}
