Commands to install required Protobuf and gRPC packages:

```bash
# This installs protoc
brew install protobuf

# This installs buf (build tool)
brew install bufbuild/buf/buf

# This installs additional binaries required for using protoc to generate Go code
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# This installs additional binariy required for using gRPC Gateway
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest

# This installs additional binariy required for using OpenAPI (Swagger) generation (not used)
go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
```
