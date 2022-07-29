# go-microservices-boilerplate
Microservices boilerplate built using Hexagonal architecture, DDD principles &amp; SOLID.

It was built in a Monorepo structure where all the services are part of a big repository, even though they're 
still running independently.


## Dependencies Packages

### Mocks
#### Ginkgo: Test Framework
Ref: https://github.com/onsi/ginkgo/v2
#### Gomega: Assertion
Ref: https://github.com/onsi/gomega
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