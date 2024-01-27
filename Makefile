SRC_DIR := $(shell git rev-parse --show-toplevel)
PKG_DIR := $(SRC_DIR)/pkg
GRPC_GEN_DIR := $(PKG_DIR)/gen
CMD_DIR := $(SRC_DIR)/cmd
BUILD_DIR := $(SRC_DIR)/build
CUR_DIR := $(shell pwd)
MODEL_DIR := $(SRC_DIR)/model

all: build

codegen:
	@echo "Generating GRPC"
	rm -rf $(GRPC_GEN_DIR)
	buf generate $(MODEL_DIR)

build: codegen lint
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