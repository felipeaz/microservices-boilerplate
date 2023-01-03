package tools

import "github.com/gin-gonic/gin"

func Register(router *gin.Engine) {
	httpRouteSwagger(router)
	httpRouteMetrics(router)
}
