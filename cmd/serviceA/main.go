package main

import (
	"log"

	"app/build/config"
	"app/build/env"
	"app/build/flags"
	"app/init/server"
	"app/internal/serviceA/handler"
	"app/internal/serviceA/repository"
	"app/internal/serviceA/service"

	_ "app/api/docs"
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
