.PHONY: server client

all: deps vet test format

deps:
	go get -v google.golang.org/grpc
	go get -v github.com/golang/protobuf/protoc-gen-go

vet:
	go vet

test:
	go test -v ./...

format:
	go fmt ./...

server:
	go run server/main.go

client:
	go run client/main.go
