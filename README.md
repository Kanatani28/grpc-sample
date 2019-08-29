# grpc-sample
gRPC sample project

## proto file complie
```
protoc --proto_path ./proto --go_out=plugins=grpc:./pb increment.proto
```

## dependency
```
go get -u google.golang.org/grpc
go get -u github.com/golang/protobuf/protoc-gen-go
```

## start server
```
go run ./server
```

## start client
```
go run ./client
```
