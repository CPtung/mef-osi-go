# Prerequisites
Go, any one of the three latest major [releases of Go](https://golang.org/doc/devel/release.html).

For installation instructions, see Go’s Getting Started guide.

[Protocol buffer](https://developers.google.com/protocol-buffers) compiler, protoc, [version 3](https://protobuf.dev/programming-guides/proto3).

For installation instructions, see [Protocol Buffer Compiler Installation](https://grpc.io/docs/protoc-installation/).

- Go plugins for the protocol compiler:
    1. Install the protocol compiler plugins for Go using the following commands:
        ```
        $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
        $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
        ```
    2. Update your PATH so that the protoc compiler can find the plugins:
        ```
        $ export PATH="$PATH:$(go env GOPATH)/bin"
        ```

# Generate gRPC code
Before you can use the new OSI service method, you need to recompile the updated .proto file.

For example, in the rpc/serial directory, run the following command:
```
$ protoc --proto_path=. --go-grpc_out=. \
    --go-grpc_opt=require_unimplemented_servers=false \
    ./serial.proto
```

This will regenerate the serial/serial.pb.go and serial/serial_grpc.pb.go files, which contain:

Code for populating, serializing, and retrieving `Serial` message types.
Generated client and server code.
