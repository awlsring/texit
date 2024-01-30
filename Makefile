SRC_DIR := $(shell git rev-parse --show-toplevel)
PKG_DIR := $(SRC_DIR)/pkg
GEN_DIR := $(PKG_DIR)/gen
CMD_DIR := $(SRC_DIR)/cmd
BUILD_DIR := $(SRC_DIR)/build
CUR_DIR := $(shell pwd)
MODEL_DIR := $(SRC_DIR)/model
HEADSCALE_22_MODEL := $(MODEL_DIR)/headscale/headscale-v0.22.3.json
HEADSCALE_22_GEN := $(GEN_DIR)/headscale/v0.22.3

all: build

test:
	go test -v ./...

vet:
	go vet ./...

codegen:
	@echo "Cleaning up"
	rm -rf $(GEN_DIR)
	@echo "Generating GRPC"
	buf generate $(MODEL_DIR)
	@echo "Generating Headscale Client for v0.22.3"
	mkdir -p $(HEADSCALE_22_GEN)
	swagger generate client -f $(HEADSCALE_22_MODEL) -A headscale -t $(HEADSCALE_22_GEN)

build: codegen vet lint test
	@echo "Tidying up"
	go mod tidy
	@echo "Building api"
	go build -o $(BUILD_DIR)/api $(CMD_DIR)/api/main.go
	@echo "Building cli"
	go build -o $(BUILD_DIR)/cli $(CMD_DIR)/cli/main.go

lint:
	@echo "Linting go files"
	golangci-lint run --fix
	@echo "Linting proto files"
	cd $(MODEL_DIR) && buf lint

.PHONY: all