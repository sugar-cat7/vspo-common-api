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

wire_gen: export PATH := $(PATH):$(GOBIN)
wire_gen:
	go generate ./app/di/wire.go

mock_gen: export PATH := $(PATH):$(GOBIN)
mock_gen:
	go generate ./domain/...

swagger_gen: export PATH := $(PATH):$(GOBIN)
swagger_gen:
	swag init

lint:
	golangci-lint run ./...

firestore_emulator:
	gcloud beta emulators firestore start --host-port=0.0.0.0:8081

firestore_emulator_clear:
	gcloud beta emulators firestore clear

run:
	go run main.go
