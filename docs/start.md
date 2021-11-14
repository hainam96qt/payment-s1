# Config to run
### 1. ENV variable:
```
export MYSQLDB="mysqldb:mysqldb@tcp(localhost:3306)/mysqldb?charset=utf8mb4&parseTime=True&loc=Local"
```
### 2. Generate:
sqlc: 
```sh
$ sqlc generate
```
protoc: 
```sh
$ ./script/protoc-gen.sh
```

### 3.Docker:
```sh
$ ./docker/docker-compose up -d
```

### 4 .Run:
```sh
$ go install ./...
$ controller
$ worker
```
 must not be null
Curl example:
```
curl --request POST \
  --url http://localhost:8082/MasterPlaylist \
  --header 'content-type: application/json' \
  --data '{
	"method":"CheckMasterPlaylistFormat",
	"resource_id":"0211edf1-1e3e-4b7d-9593-2b8929d8d3c7"
}	'
```
