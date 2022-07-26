package main

import (
	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/serviceA/api"
	"microservices-boilerplate/internal/serviceA/handler"
	"microservices-boilerplate/internal/serviceA/repository"
	"microservices-boilerplate/internal/serviceA/service"
)

func main() {
	cfg := config.Build(
		env.Build(),
		flags.Build(),
	)
	apiServer := server.New(
		api.New(
			handler.New(
				service.New(cfg.Log, repository.New(cfg.Database, cfg.Cache)),
			),
			middleware.New(),
		),
	)
	apiServer.Run(cfg.Addr)
}
