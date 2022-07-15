p-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    users/users.proto

tidy:
	@go mod tidy

run-server:
	@go run server/main.go

proto: p-user

build-client:
	@go build -o client client/main.go
	@mv client/main app-client

build-server:
	@go build -o server server/main.go 
	@mv server/main app-server

.PHONY: proto p-user tidy run-client run-server build-client build-server