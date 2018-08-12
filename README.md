# ogay gRPC Example

## Project Setup

```bash
$ git clone git@github.com:spilth/ogay-grpc.git
$ cd ogay-grpc
$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/protoc-gen-go
```

### Running the Server

```bash
$ go run server/main.go
```

### Running the Client

```bash
$ go run client/main.go
2018/08/12 14:58:24 'Hello' translates to 'Ellohay'
$ go run client/main.go Go
2018/08/12 14:57:51 'Go' translates to 'Ogay'
```

## Regenerating the Generated Code

```bash
$ protoc ogay.proto --go_out=plugins=grpc:.
```
