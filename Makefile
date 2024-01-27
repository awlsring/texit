SRC_DIR := $(shell git rev-parse --show-toplevel)
PKG_DIR := $(SRC_DIR)/pkg
GRPC_GEN_DIR := $(PKG_DIR)/gen/client-grpc/v1
APP_DIR := $(SRC_DIR)/cmd/api
APP_PKG_DIR := $(SRC_DIR)/internal/app/api
BUILD_DIR := $(SRC_DIR)/build
CUR_DIR := $(shell pwd)
MODEL_DIR := $(SRC_DIR)/model/api

all: build

codegen:
	@echo "Generating GRPC"
	rm -rf $(GRPC_GEN_DIR)
	buf generate $(MODEL_DIR)

build: lint
	@echo "Building project"
	go mod tidy
	go build -o $(BUILD_DIR)/api $(APP_DIR)/main.go

lint:
	golangci-lint run --fix --timeout 10m
	cd $(MODEL_DIR) && go run github.com/bufbuild/buf/cmd/buf lint

.PHONY: all