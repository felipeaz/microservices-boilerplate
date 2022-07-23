package api

import "github.com/gin-gonic/gin"

type Api interface {
	RegisterRoutes(router *gin.Engine)
}
