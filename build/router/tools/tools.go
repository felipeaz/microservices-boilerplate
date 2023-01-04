package tools

import "github.com/gin-gonic/gin"

func RegisterStandardTools(router *gin.Engine) {
	httpRouteSwagger(router)
}
