package main

import (
	"context"
	"fmt"
	"os"

	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli"
	"github.com/awlsring/texit/internal/app/ui/adapters/primary/cli/handler"
	api_gateway "github.com/awlsring/texit/internal/app/ui/adapters/secondary/gateway/api"
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

func initClient(address string, key string) texit.Invoker {
	c, err := texit.NewClient(address, Sec{key: key})
	panicOnErr(err)
	return c
}

func isInitCommand(args []string) bool {
	return len(args) > 1 && args[1] == "init"
}

func main() {
	if isInitCommand(os.Args) {
		cli.InitDefaultConfig()
		fmt.Println("Texit default config has been initialized. Please edit the file at ~/.texit/config.yaml to set your server address and api key.")
		return
	}

	cli.MakeTexitDir()
	cfg, err := cli.LoadConfig()
	panicOnErr(err)
	client := initClient(cfg.Api.Address, cfg.Api.ApiKey)
	apiGw := api_gateway.New(client)
	svc := api.NewService(apiGw)
	hdl := handler.New(svc)
	tool := cli.New(hdl)
	panicOnErr(tool.Run())
}
