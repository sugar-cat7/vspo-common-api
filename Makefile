.PHONY: build test wire_gen mock_gen install-tools

GOBIN := $(shell go env GOPATH)/bin
MOCKGEN := $(GOBIN)/mockgen
WIRE := $(GOBIN)/wire

install-tools:
	go get github.com/golang/mock/mockgen@v1.6.0
	go get github.com/google/wire/cmd/wire@v0.5.0

build:
	go build ./...

test:
	go test ./...

wire_gen:
	go generate ./app/di/wire.go

mock_gen:
	export PATH=$PATH:$(go env GOPATH)/bin
	go generate ./domain/...

lint:
	golangci-lint run ./...
