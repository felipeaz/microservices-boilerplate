package router

import (
	"github.com/gin-gonic/gin"

	"app/api/middleware"
	"app/build/router/tools"
)

func New(routerEngine *gin.Engine) *gin.Engine {
	registerStandardMiddlewares(routerEngine)
	tools.RegisterStandardTools(routerEngine)

	return routerEngine
}

func registerStandardMiddlewares(router *gin.Engine) {
	corsMiddleware := middleware.NewCorsMiddleware()
	prometheusMiddleware := middleware.NewPrometheusMiddleware(router)

	router.Use(corsMiddleware.HandleFunc())
	router.Use(prometheusMiddleware.HandleFunc())
}
