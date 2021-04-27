# go-app

Quick start [Go-gRPC](https://grpc.io/docs/languages/go/quickstart/)

## Run

Recompile the updated .proto file:
```console
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative model/sample.proto
```

Run unit test:
```console
$ go test
```

Run the server:
```console
$ go build
$ ./go-app daemon
```

From another terminal, run the client:
```console
$ go build
$ ./go-app grpchello
```

From another terminal, call API:
```console
$ ./go-app grpc-order-create https://abc.xyz
$ ./go-app grpc-order-read 1
$ curl localhost:9000/orders/1
```
