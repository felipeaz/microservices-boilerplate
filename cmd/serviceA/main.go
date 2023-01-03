package main

import (
	"log"

	"microservices-boilerplate/build/config"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/serviceA/handler"
	"microservices-boilerplate/internal/serviceA/repository"
	"microservices-boilerplate/internal/serviceA/service"

	_ "microservices-boilerplate/api/docs"
)

// @title       Service A Swagger Example API
// @version     1.0
// @description This is a sample server.

// @host     localhost:8085
// @BasePath /api/v1

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered in main")
		}
	}()

	cfg := config.Build(env.Build(), flags.Build())

	server.New(
		handler.New(
			&handler.DependenciesNode{
				Service: service.New(
					&service.DependenciesNode{
						Log: cfg.Logger,
						Repository: repository.New(
							&repository.DependenciesNode{
								Database: cfg.Database,
								Cache:    cfg.Cache,
							},
						),
					},
				),
				Router: cfg.Router,
			},
		),
	).Run(cfg.Port)
}
