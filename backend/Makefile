# Makefile for Backend workflow (buf + go)

ifeq ($(OS),Windows_NT)
	IS_WINDOWS := 1
endif

# repo layout (this Makefile is in e:\Atlas\backend)
API_DIR := api

.PHONY: help env init cli plugin dep lint gen tidy all

# show environment variables
env:
	@echo "API_DIR: $(API_DIR)"
	@echo "IS_WINDOWS: $(IS_WINDOWS)"

# initialize develop environment
init: cli plugin all
	@echo "init done."

# install/update cli tools
cli:
	@go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install github.com/google/wire/cmd/wire@latest

# install/update protoc plugins used by buf.gen.yaml
plugin:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest

# default workflow
all:
	@echo "all done."

# show help
help:
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  env    - show variables"
	@echo "  init   - cli + plugin + all"
	@echo "  cli    - install/update buf + wire"
	@echo "  plugin - install/update protoc-gen-*"
	@echo "  dep    - go mod download"
	@echo "  lint   - golangci-lint run"
	@echo "  gen    - (placeholder) generate done"
	@echo "  tidy   - go mod tidy"
	@echo "  all    - default workflow"
	@echo ""

.DEFAULT_GOAL := help
