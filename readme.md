# How to run:
Build Dockerfile from server/: ```docker build . -t grpc_redis_server:1.0.0```
<br/>
Run the server/: ```docker compose up```
<br/>
Install the client/: ```go build && go install```
<br/>
Make sure you have added your go install bin directory to PATH
<br/>
Run the client: ```go_client list``` or ```go_client add <key> <value>```