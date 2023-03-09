Commands to install required Protobuf and gRPC packages:

```bash
# This installs protoc
brew install protobuf

# This install additional binaries required for using protoc to generate Go code
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
