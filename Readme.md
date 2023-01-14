# Introduction to gRPC for MicroServices

Build Proto :

Golang 
```
protoc --go-grpc_out=. --go-grpc_opt=require_unimplemented_servers=false --go_out=. proto/*.proto
```