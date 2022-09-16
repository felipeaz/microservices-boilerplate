package main

import (
	"github.com/gin-gonic/gin"
	"log"

	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/api/serviceB"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/serviceB/handler"
	"microservices-boilerplate/internal/serviceB/repository"
	"microservices-boilerplate/internal/serviceB/service"

	_ "microservices-boilerplate/api/docs"
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

	cfg := config.Build(env.Build(), flags.Build())
	router := gin.Default()
	router.Use(middleware.New().Cors())

	api := server.New(
		serviceB.NewApi(
			handler.New(
				&handler.Config{
					Service: service.New(
						&service.Config{
							Log: cfg.Logger,
							Repository: repository.New(
								&repository.Config{
									Database: cfg.Database,
									Cache:    cfg.Cache,
								},
							),
						},
					),
				},
			),
			router,
		),
	)

	api.Run(cfg.Port)
}
