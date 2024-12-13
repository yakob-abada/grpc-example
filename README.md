# GRPC in Go
This is a comprehensive example of building a gRPC service in Go, implementing the 
[AIP-158](https://google.aip.dev/158) Google standard for pagination. 
The project follows the Domain-Driven Design (DDD) pattern, emphasizing separation of concerns and maintainability. 
Additionally, the application has been containerized using Docker to ensure consistency and ease of deployment.

## Installation

* copy `env.example` file to `.env` `cp env.example .env`
* Build and start the application with Docker Compose `docker-compose up --build`
* You migration up inside the container and could be done by the following.
* Access the container `docker exec -it grpc-server bash` 
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