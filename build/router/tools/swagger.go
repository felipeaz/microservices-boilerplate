package tools

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	defaultSwaggerPath = "/swagger/*any"
)

func httpRouteSwagger(router *gin.Engine) {
	router.GET(defaultSwaggerPath, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
