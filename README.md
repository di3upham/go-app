# go-app

## Run

Recompile the updated .proto file:
```console
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative model/sample.proto
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
