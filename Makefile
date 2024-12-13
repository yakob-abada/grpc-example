build_proto:
	protoc --go_out=./explore --go_opt=paths=source_relative \
    --go-grpc_out=./explore --go-grpc_opt=paths=source_relative \
    ./proto/*.proto --experimental_allow_proto3_optional

unit_test:
	go test ./pkg/...

integration_test:
	go run db/fixtures/main.go
	go test ./integration/...

migration_up:
	go run db/create_tables/main.go

migration_down:
	go run db/drop_tables/main.go

fixtures:
	go run db/fixtures/main.go