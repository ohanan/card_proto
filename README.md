# card_proto

## Proto buffer
using [buf](https://buf.build/docs/introduction)
or use 
```shell
go run ./cmd/prepare
protoc \
  --go_out=pkg/protoservice \
  --go_opt=paths=source_relative \
  --go-grpc_out=pkg/protoservice \
  --go-grpc_opt=paths=source_relative \
  --go-grpc_opt=require_unimplemented_servers=false \
  --card-proto-plugin_out=pkg/protoservice \
  --card-proto-plugin_opt=paths=source_relative \
  ./**/*.proto
```
