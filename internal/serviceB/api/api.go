package api

import (
	"github.com/gin-gonic/gin"

	_api "microservices-boilerplate/api"
	"microservices-boilerplate/internal/serviceB/handler"
)

type api struct {
	handler handler.Handler
}

func New(h handler.Handler) _api.Api {
	return &api{
		handler: h,
	}
}

func (a *api) RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	vGroup := apiGroup.Group("/v1")
	vGroup.GET("/item", a.handler.Get)
	vGroup.GET("/item/:id", a.handler.Find)
	vGroup.POST("/item", a.handler.Create)
	vGroup.PUT("/item/:id", a.handler.Update)
	vGroup.DELETE("/item/:id", a.handler.Delete)
}
