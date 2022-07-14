p-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    users/users.proto

tidy:
	@go mod tidy

run-server:
	@go run server/main.go

proto: p-user

.PHONY: proto p-user tidy run-client run-server