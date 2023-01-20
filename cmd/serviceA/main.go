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

	"github.com/gin-gonic/gin"
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

	cfg := config.Build(
		config.BuildArgs{
			Env:    env.Build(),
			Flags:  flags.Build(),
			Router: gin.Default(),
		},
	)

	repo := repository.New(
		&repository.DependenciesNode{
			Database: cfg.Database,
			Cache:    cfg.Cache,
		},
	)

	api := service.New(
		&service.DependenciesNode{
			Log:        cfg.Logger,
			Repository: repo,
		},
	)

	handlerGateway := handler.New(
		&handler.DependenciesNode{
			Service: api,
			Router:  cfg.Router,
		},
	)

	err := server.New(handlerGateway.GetRouter()).Run(cfg.ServicePort)
	if err != nil {
		log.Fatal(err)
	}
}
