.PHONY: build
build:
	go build -v ./cmd/apiserver


.PHONY: grpc
grpc:
	go build -v ./cmd/grpcserver

.DEFAULT_GOAL := build