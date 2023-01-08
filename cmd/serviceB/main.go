package main

import (
	"log"

	"app/build/config"
	"app/build/env"
	"app/build/flags"
	"app/init/server"
	"app/internal/serviceB/handler"
	"app/internal/serviceB/repository"
	"app/internal/serviceB/service"

	_ "app/api/docs"

	"github.com/gin-gonic/gin"
)

// @title       Service B Swagger Example API
// @version     1.0
// @description This is a sample server.

// @host     localhost:8086
// @BasePath /api/v1

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered in main")
		}
	}()

	cfg := config.Build(
		config.BuildArgs{
			Env:    env.Build(),
			Flags:  flags.Build(),
			Router: gin.Default(),
		},
	)

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
	).Run(cfg.ServicePort)
}
