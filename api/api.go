package api

import "github.com/gin-gonic/gin"

type Api interface {
	RegisterRoutes()
	GetRouter() *gin.Engine
}
