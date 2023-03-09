#!/bin/sh

# protoc --go_out=. pb/test_obj.proto

protoc -I=pkg/pb \
  --go_out=paths=source_relative:pkg/pb \
  --go-grpc_out=paths=source_relative:pkg/pb \
  pkg/pb/*.proto
