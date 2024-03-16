# Description: Makefile for server

BIN_DIR := bin
SERVER_BIN_NAME := server

.PHONY: build
build:
	CGO_ENABLED=0 go build -o $(BIN_DIR)/$(SERVER_BIN_NAME) cmd/api-server/main.go

.PHONY: linux-build
linux-build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(SERVER_BIN_NAME) cmd/api-server/main.go

.PHONY: run
run:
	go run cmd/api-server/main.go

.PHONY: test
test:
	go test -v ./...
