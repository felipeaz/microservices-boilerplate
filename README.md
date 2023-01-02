# microservices-boilerplate
Microservices boilerplate built using Hexagonal architecture, DDD principles &amp; SOLID.

It was built in a Monorepo structure where all the services are part of a big repository, even though they're 
still running independently.

## Application Guide
This project has been automated with Makefile to simplify the configuration and how to run the services. To execute the
functions bellow, it's required that you CD at the project root folder -> `microservices-boilerplate`.

### Installation
The application depends on `Kong API Gateway` to manage the microservices. It's important to configure everything
before executing the API itself. Spoiler: Kong will take care of the service discovery, defining routes and set a JWT  
authentication method.

To install and configure the app, run `make install`.

### Running API
This command will start the API by executing the docker-compose files. Make sure you have installed the application
before executing the running step.

To execute the app, run `make run`.

You can run `make stop` to stop the application and shut down the docker containers.

### Running Tests
This command executes all test cases in coverage mode and generates an HTML page with the output. The files generated 
with this command will be at `test/coverage`.

Run `make tests` to execute all tests.

## Dependencies Packages

### Tests
#### Ginkgo: Test Framework
Ref: https://github.com/onsi/ginkgo/v2
#### Gomega: Assertion
Ref: https://github.com/onsi/gomega

### Mocks
#### Testify: Mocks implementation
Ref: https://github.com/stretchr/testify
#### Vektra Mockery: Mocks Generator
Ref: https://github.com/vektra/mockery

### GORM
#### Entities & DB operations
Ref: https://gorm.io/

### Redis
#### Cache management
Ref: https://github.com/gomodule/redigo/redis

### UUID
#### Secure ID generator
Ref: https://github.com/satori/go.uuid

### Cors Middleware
#### Enable API to external origins
Ref: https://github.com/itsjamie/gin-cors

### Observability
#### Application monitoring using Prometheus and Grafana
Ref: https://grafana.com/
Ref: https://prometheus.io/

### API Documentation
#### Gin-Swagger: Swagger API Definitions
Ref: https://github.com/swaggo/gin-swagger
#### Swaggo Files: Doc files manager
Ref: https://github.com/swaggo/files
#### Swaggo: Doc Generator
Ref: https://github.com/swaggo/swag

#### API Doc: ${service_host}/swagger/index.html

## Application High Level Architecture
![Microservices Boilerplate drawio (1)](https://user-images.githubusercontent.com/32846823/182005597-e9512985-27d9-45ce-b74f-6b0bd4e8f9f2.png)
