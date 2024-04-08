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

# Create OSI service
Now, create a Go file under `interface/{MODULE_NAME}` to implement your gRPC service:
```go
// interface/serial/serial.go
package serial

import (
	"context"

	rpc "github.com/MOXA-IPC/mef-osi-go/rpc/serial"
)

type SerialImpl struct {
}

func NewService() rpc.SerialServer {
	return &SerialImpl{}
}

func (s *SerialImpl) GetSerial(ctx context.Context, empty *rpc.SerialEmptyRequest) (*rpc.SerialReply, error) {
	return nil, nil
}

func (s *SerialImpl) SetSerial(ctx context.Context, request *rpc.SerialRequest) (*rpc.SerialReply, error) {
	return nil, nil
}

```

# Register OSI service
After the gRPC service code is ready, it has to register to the server.

```go
// pkg/server/server.go

...

import (
    serial "github.com/MOXA-IPC/mef-osi-go/interface/serial"
	rpc_serial "github.com/MOXA-IPC/mef-osi-go/rpc/serial"
)

func New() *Server {
	if err := os.MkdirAll(path.Dir(sockfd), os.ModePerm); err != nil {
		log.Printf("mkdirall error %s\n", err.Error())
		return nil
	}

	// start grpc reverse proxy socket
	syscall.Unlink(sockfd)
	sockProxy, err := net.Listen("unix", sockfd)
	if err != nil {
		log.Printf("failed to listen: %v\n", err)
		return nil
	}

	/* Register all grpc services */
	grpcServer := grpc.NewServer()
	rpc_serial.RegisterSerialServer(grpcServer, serial.NewService())
	///////////////////////////////

	return &Server{sockProxy, grpcServer}
}

func (s *Server) Serve() {
	if err := s.server.Serve(s.listener); err != nil {
		log.Printf("grpc serve error %s\n", err.Error())
	}
}

```

# Build and run your OSI service
Build and run your Go application:

```shell
$ go build -o main cmd/main.go
$ ./main
```

# Test your OSI service
Follow the naming conventions, Go’s testing package comes with an expectation that any test file must have a `_test.go` suffix. For example, if we would have a OSI file called `serial.go` its test file must be named `serial_test.go`.

Now you can create a test file under `osi_test` and assume we are going to test the `GetProfile` function of serial.
```go
...

func (s *OsiV1TestSuite) TestSerial_GetProfile() {
    // The first, we need a OSI client that's connecting to the OSI mock server.
    cli := client.New("localhost:8880")
    assert.NotNil(s.T(), cli)
    defer cli.Close()

    // Bind serial gRPC to the client
    serial := pb.NewSerialClient(cli)
    assert.NotNil(s.T(), serial)

	// Contact the server and print out its response.
    ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(3)*time.Second)
    defer cancel()
    r, err := serial.GetSerial(ctx, &pb SerialEmptyRequest{})
    if err != nil {
        log.Fatalf("could not get serial: %v", err)
    }
    log.Printf("Greeting: %v", r.GetProfiles())
}

```

As the test case done, you can use go test tool to show the result and even the test coverage.
```shell
$ go test -v -cover -coverpkg=./osi/...,./pkg/... -coverprofile=coverage.out ./osi_test
$ go tool cover -func coverage.out
```
