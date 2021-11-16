# Config to run

### 1. Generate:
sqlc: 
```sh
$ sqlc generate
```
protoc: 
```sh
$ ./script/protoc-gen.sh
```

### 2.Docker:
```sh
$ ./docker/docker-compose up -d
```

### 3 .Run:
```sh
$ go install ./...
$ payment-challenge
```
