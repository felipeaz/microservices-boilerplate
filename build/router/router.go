package router

import (
	"github.com/gin-gonic/gin"
	"microservices-boilerplate/api/middleware"
	"microservices-boilerplate/build/router/tools"
)

func New() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.New().Cors())
	tools.Register(router)

	return router
}
