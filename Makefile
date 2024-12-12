build_proto:
	protoc --go_out=./explore --go_opt=paths=source_relative \
    --go-grpc_out=./explore --go-grpc_opt=paths=source_relative \
    ./proto/*.proto

unit_test:
	go test ./...

migration:
	go run db/create_match_table.go