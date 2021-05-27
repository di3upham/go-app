# go-app

Quick start [Go-gRPC](https://grpc.io/docs/languages/go/quickstart/), [MySQL](https://www.digitalocean.com/community/tutorials/how-to-install-mysql-on-ubuntu-20-04)

## Run

Recompile the updated .proto file:
```console
$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative model/sample.proto
```

Run unit test:
```console
$ go test
```

MySQL:
```console
mysql> CREATE USER 'user'@'localhost' IDENTIFIED BY 'password';
mysql> CREATE DATABASE dbname CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
mysql> GRANT ALL PRIVILEGES ON dbname.* TO 'user'@'localhost';
mysql> CREATE TABLE `orders` (`id` varchar(20), `status` varchar(10), `created_at` bigint, `product_url` varchar(1024), UNIQUE KEY `id` (`id`));
```

PostgreSQL:
```console
$ sudo su -postgres
$ psql -c "ALTER USER username WITH PASSWORD 'password'"
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
