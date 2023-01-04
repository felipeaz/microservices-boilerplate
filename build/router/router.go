package router

import (
	"github.com/gin-gonic/gin"

	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/build/router/tools"
)

func New() *gin.Engine {
	router := gin.Default()

	registerStandardMiddlewares(router)
	tools.RegisterStandardTools(router)

	return router
}

func registerStandardMiddlewares(router *gin.Engine) {
	corsMiddleware := middleware.NewCorsMiddleware()
	prometheusMiddleware := middleware.NewPrometheusMiddleware(router)

	router.Use(corsMiddleware.HandleFunc())
	router.Use(prometheusMiddleware.HandleFunc())
}
