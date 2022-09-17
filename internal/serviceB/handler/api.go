package handler

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"microservices-boilerplate/api"
)

func (h *Handler) RegisterRoutes() {
	h.config.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := h.config.Router.Group("/api")
	{
		vGroup := apiGroup.Group("/v1")
		{
			vGroup.GET("/b-items", h.Get)
			vGroup.GET("/b-items/:id", h.Find)
			vGroup.POST("/b-items", h.Create)
			vGroup.PUT("/b-items/:id", h.Update)
			vGroup.DELETE("/b-items/:id", h.Delete)
		}
	}
}

func (h *Handler) GetRouter() api.Router {
	return h.config.Router
}
