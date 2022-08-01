package main

import (
	"github.com/gin-gonic/gin"

	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/api/serviceA"
	_ "microservices-boilerplate/api/serviceA/docs"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/serviceA/handler"
	"microservices-boilerplate/internal/serviceA/repository"
	"microservices-boilerplate/internal/serviceA/service"
)

// @title           Service A Swagger Example API
// @version         1.0
// @description     This is a sample server.

// @host      localhost:8085
// @BasePath  /api/v1

func main() {
	cfg := config.Build(env.Build(), flags.Build())
	router := gin.Default()
	router.Use(middleware.New().Cors())

	api := server.New(
		serviceA.NewApi(
			handler.New(
				service.New(
					cfg.Logger,
					repository.New(cfg.Database, cfg.Cache),
				),
			),
			router,
		),
	)

	api.Run(cfg.Port)
}
