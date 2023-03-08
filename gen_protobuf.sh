#!/bin/sh

# protoc --go_out=. pb/test_obj.proto
protoc -I=pb --go_out=paths=source_relative:pb pb/*.proto
