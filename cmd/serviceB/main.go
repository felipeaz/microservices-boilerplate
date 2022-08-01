package main

import (
	"github.com/gin-gonic/gin"

	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/api/serviceB"
	_ "microservices-boilerplate/api/serviceB/docs"
	"microservices-boilerplate/build/config"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/init/server"
	"microservices-boilerplate/internal/serviceB/handler"
	"microservices-boilerplate/internal/serviceB/repository"
	"microservices-boilerplate/internal/serviceB/service"
)

func main() {
	cfg := config.Build(env.Build(), flags.Build())
	router := gin.Default()
	router.Use(middleware.New().Cors())

	api := server.New(
		serviceB.NewApi(
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
