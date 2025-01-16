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

## Improvements
* Increase Test Coverage: Enhance the unit and integration tests to improve overall test coverage and ensure greater reliability.
* Implement Logging System: Introduce a robust logging framework to provide better traceability, 
error reporting, and debugging capabilities across the application.
* proto file(s) should be in separate repository and should be imported to avoid repeat.