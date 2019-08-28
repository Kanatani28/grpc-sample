# grpc-sample
gRPC sample project

## proto file complie
```
protoc --proto_path ./proto --go_out=plugins=grpc:./pb increment.proto
```

## start server
```
go run ./server
```

## start client
```
go run ./client
```
