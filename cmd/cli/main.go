package main

import (
	"context"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/handler"
	api_gateway "github.com/awlsring/texit/internal/app/ui/adapters/secondary/gateway/api"
	"github.com/awlsring/texit/internal/app/ui/config"
	"github.com/awlsring/texit/internal/app/ui/core/service/api"
	"github.com/awlsring/texit/pkg/gen/texit"
)

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}

type Sec struct {
	key string
}

func (s Sec) SmithyAPIHttpApiKeyAuth(ctx context.Context, operationName string) (texit.SmithyAPIHttpApiKeyAuth, error) {
	return texit.SmithyAPIHttpApiKeyAuth{
		APIKey: s.key,
	}, nil
}

func initClient(address string) texit.Invoker {
	c, err := texit.NewClient(address, Sec{key: "changeme"})
	panicOnErr(err)
	return c
}

func main() {
	cfg, err := config.LoadFromFile("cli-config.yaml")
	panicOnErr(err)
	client := initClient(cfg.Api.Address)

	apiGw := api_gateway.New(client)
	svc := api.NewService(apiGw)
	hdl := handler.New(svc)
	tool := cli.New(hdl)
	panicOnErr(tool.Run())
}
