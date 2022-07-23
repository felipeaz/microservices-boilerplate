package main

import (
	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/serviceB/api"
	"microservices-boilerplate/internal/serviceB/handler"
	"microservices-boilerplate/internal/serviceB/repository"
	"microservices-boilerplate/internal/serviceB/service"
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
