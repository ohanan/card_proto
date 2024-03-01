#!/usr/bin/env bash
rm -rf ./pkg/protoservice/proto
rm ./pkg/protoservice/*.pb.go
go run ./cmd/prepare
go install ./cmd/protoc-gen-card-proto-plugin
buf format -w
protoc \
  --go_out=pkg/protoservice \
  --go_opt=paths=source_relative \
  --go-grpc_out=pkg/protoservice \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  --card-proto-plugin_out=pkg/protoservice \
  --card-proto-plugin_opt=paths=source_relative \
  service.proto ./proto/*.proto