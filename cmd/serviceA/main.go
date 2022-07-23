package main

import (
	"microservices-boilerplate/build/flags"
	"os"

	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/serviceA/api"
	"microservices-boilerplate/internal/serviceA/handler"
	"microservices-boilerplate/internal/serviceA/repository"
	"microservices-boilerplate/internal/serviceA/service"
	"microservices-boilerplate/internal/storage/redis"
)

func main() {
	cfg := flags.Build()

	// db := redis.New(os.Getenv("DB_HOST"))
	cache := redis.New(os.Getenv("REDIS_HOST"))
	logger := pkg.NewLogger(*cfg.Debug)

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
