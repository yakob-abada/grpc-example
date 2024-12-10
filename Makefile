build_proto:
	protoc --go_out=./explore --go_opt=paths=source_relative \
    --go-grpc_out=./explore --go-grpc_opt=paths=source_relative \
    https://github.com/muzzapp/backend-interview-task/blob/main/explore-service.proto

unit_test:
	go test ./...