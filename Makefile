go_container := grpc-server
docker := docker-compose
compose := $(docker) --file docker-compose.yaml
docker_exec := docker exec
args = $(filter-out $@,$(MAKECMDGOALS))

build:
	$(docker) up --build -d

up:
	$(docker) up -d

container_access:
	$(docker_exec) -it $(go_container) bash

build_proto:
	$(docker_exec) $(go_container) protoc --go_out=./explore --go_opt=paths=source_relative \
    --go-grpc_out=./explore --go-grpc_opt=paths=source_relative \
    ./proto/*.proto --experimental_allow_proto3_optional

unit_test:
	$(docker_exec) $(go_container) go test ./pkg/...

integration_test:
	$(docker_exec) $(go_container) go run db/fixtures/main.go && go test ./integration/...

migration_up:
	$(docker_exec) $(go_container) go run db/create_tables/main.go

migration_down:
	$(docker_exec) $(go_container) go run db/drop_tables/main.go

fixtures:
	$(docker_exec) $(go_container) go run db/fixtures/main.go