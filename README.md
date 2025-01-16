# GRPC in Go
This repository provides a gRPC example project that demonstrates the implementation of [AIP-158](https://google.aip.dev/158) pagination 
standards and adherence to Domain-Driven Design (DDD) principles. 
The project is containerized using Docker, making it easy to set up and run in a consistent environment.


## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/doc/install) (if running locally without Docker)

### Clone the Repository

```
git clone https://github.com/yakob-abada/grpc-example.git
cd grpc-example
```

* copy `env.example` file to `.env` `cp env.example .env`
* Build and start the application with Docker Compose `make build`
* You migration up inside the container and could be done by the following.
* Access the container `make container_access` 
* To apply migration_up `make migration_up`
* To execute unit tests `make unit_test`
* To execute integration tests `make unit_test`
* To apply migration_down `make migration_down`
* To run fixtures `make fixtures`

* phpmyadmin container is run so database can be accessed from a browser http://localhost:8081

#### Using `grpcurl`

[grpcurl](https://github.com/fullstorydev/grpcurl) is a command-line tool to interact with gRPC services.

1. Install `grpcurl`.
2. Call the service:

    ```bash
    grpcurl -plaintext localhost:50051 <ServiceName>/<MethodName>
    ```

#### Using a gRPC Client

Use any gRPC client (e.g., Postman or a custom client) to connect to `localhost:50051`.

### Protobuf Definitions

The `.proto` files are located in the `proto` directory. They define the gRPC service, messages, and pagination structure. Ensure you compile the `.proto` files to generate the necessary language bindings for your client. Use the `Makefile` to simplify the process:

## Project Structure

- `cmd/`: Entry point of the application.
- `pkg/`: Application logic adhering to DDD principles.
- `proto/`: Protobuf definitions.
- `internal/`: Internal utilities and configurations.
- `Dockerfile`: Docker setup for containerization.
- `README.md`: Project documentation.

## Improvements
* Increase Test Coverage: Enhance the unit and integration tests to improve overall test coverage and ensure greater reliability.
* Implement Logging System: Introduce a robust logging framework to provide better traceability, 
error reporting, and debugging capabilities across the application.
* proto file(s) should be in separate repository and should be imported to avoid repeat.