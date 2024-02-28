# card_proto

## Proto buffer
using [buf](https://buf.build/docs/introduction)
or use 
```shell
protoc --go_out=pkg/proto --go_opt=paths=source_relative --go-grpc_out=pkg/proto --go-grpc_opt=paths=source_relative --go-grpc_opt=require_unimplemented_servers=false --card-proto-plugin_out=pkg/proto --card-proto-plugin_opt=paths=source_relative card.proto 
```
