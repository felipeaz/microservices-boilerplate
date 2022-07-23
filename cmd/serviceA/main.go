package main

import (
	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/serviceA/api"
	"microservices-boilerplate/internal/serviceA/handler"
	"microservices-boilerplate/internal/serviceA/repository"
	"microservices-boilerplate/internal/serviceA/service"
	"microservices-boilerplate/internal/storage/redis"
)

func main() {
	cfg := config.Build()

	// db := postgresql.New(config.Env.DBHost)
	cache := redis.New(cfg.Env.CacheHost)
	logger := pkg.NewLogger(*cfg.Flags.Debug)

	apiServer := server.New(
		api.New(
			handler.New(
				service.New(logger, repository.New(cache)),
			),
			middleware.New(),
		),
	)

	apiServer.Run(cfg.Env.Host)
}
