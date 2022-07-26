package main

import (
	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/serviceB/api"
	"microservices-boilerplate/internal/serviceB/handler"
	"microservices-boilerplate/internal/serviceB/repository"
	"microservices-boilerplate/internal/serviceB/service"
)

func main() {
	cfg := config.Build(
		env.Build(),
		flags.Build(),
	)
	apiServer := server.New(
		api.New(
			handler.New(
				service.New(
					cfg.Logger,
					repository.New(cfg.Database, cfg.Cache),
				),
			),
			middleware.New(),
		),
	)
	apiServer.Run(cfg.Port)
}
