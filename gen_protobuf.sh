#!/bin/sh

protoc -I=proto \
  --go_out=paths=source_relative:pkg/pb \
  --go-grpc_out=paths=source_relative:pkg/pb \
  proto/*.proto
