# test-grpc
Just to try out gRPC in Go

# Setup
- Install Go
- Install proto compiler plugin for Go and grpc
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
  ```
- Set the binary installed directory in path for `protoc` to find it.
  Run `export PATH="$PATH:$(go env GOPATH)/bin" or add it to `~/.bash_profile`
- To test, run server by `go run server.go` and client by `go run client.go`
- Warning: the module contains two main methods in main package i.e.
  one in `server.go` and other in `client.go`, which might give warning but
  should be fine as this is a test project and client and server are running
  seperately.

# Notes
- gRPC service and request response format are defined in `protos/service.proto`.
- `*.proto` file is used to generate request response struct for Go client and
  server along with interface for service rpc methods in `*.pb.go` and `*_grpc.pb.go`
  file respectively.
- To generate the Go files:
  ```bash
  protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative protos/service.proto
  ```
- `--go_out` tells proto compiler to generate Go structs (`*.pb.go`) which
  depend on plugin `protoc-gen-go`.
- `--go-grpc_out` tells the `protoc` to generate Go client code and service
  interface (`*_grpc.pb.go`) using another plugin `protoc-gen-go-grpc`.
- Once the code is generated, server can implement the service interface using
  the structs defined in `.pb.go` file and start to listen on a port.
- Similarly, client can connect to the server using client defined in `_grpc.pb.go`
  file to call the service methods and use the request and response structs
  defined in `.pb.go` files.

# References
- https://grpc.io/docs/languages/go/quickstart/
- https://grpc.io/docs/languages/go/basics/
