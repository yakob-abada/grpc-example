# GRPC in Go
This is a comprehensive example of building a gRPC service in Go, implementing the 
[AIP-158](https://google.aip.dev/158) Google standard for pagination. 
The project follows the Domain-Driven Design (DDD) pattern, emphasizing separation of concerns and maintainability. 
Additionally, the application has been containerized using Docker to ensure consistency and ease of deployment.

## Installation

* copy `env.example` file to `.env` `cp env.example .env`
* Build and start the application with Docker Compose `docker-compose up --build`
* Access the container `docker exec -it grpc-server bash` 
  * To execute unit tests `make unit_test`
  * To apply migrations `make migration`

* phpmyadmin container is run so database can be accessed from a browser http://localhost:8081

## Improvements
* Increase Test Coverage: Enhance the unit and integration tests to improve overall test coverage and ensure greater reliability.
* Implement Logging System: Introduce a robust logging framework to provide better traceability, 
error reporting, and debugging capabilities across the application.