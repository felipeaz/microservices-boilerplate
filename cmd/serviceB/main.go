package main

import (
	"os"

	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/serviceB/api"
	"microservices-boilerplate/internal/serviceB/handler"
	"microservices-boilerplate/internal/serviceB/repository"
	"microservices-boilerplate/internal/serviceB/service"
	"microservices-boilerplate/internal/storage/redis"
)

func main() {
	// db := redis.New(os.Getenv("DB_HOST"))
	cache := redis.New(os.Getenv("REDIS_HOST"))
	logger := pkg.NewLogger(false)

	apiServer := server.New(
		api.New(
			handler.New(
				service.New(logger, repository.New(cache)),
			),
			middleware.New(),
		),
	)

	apiServer.Run(os.Getenv("SERVER_ADDR"))
}
