# Description: Makefile for server

BIN_DIR := bin

.PHONY: go-build
go-build:
	go build -o $(BIN_DIR)/server cmd/api-server/main.go

.PHONY: go-run
go-run:
	go run cmd/api-server/main.go

.PHONY: go-test
go-test:
	go test -v ./...
