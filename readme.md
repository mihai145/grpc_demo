# How to run:
Build Dockerfile from server/: ```docker build . -t grpc_redis_server:1.0.0```
<br/>
Run the server: ```docker compose up```
<br/>
Run the client: ```go run client.go list``` or ```go run client.go add <key> <value>```