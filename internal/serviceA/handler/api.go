package handler

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) RegisterRoutes() {
	h.config.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := h.config.Router.Group("/api")
	{
		vGroup := apiGroup.Group("/v1")
		{
			vGroup.GET("/a-items", h.Get)
			vGroup.GET("/a-items/:id", h.Find)
			vGroup.POST("/a-items", h.Create)
			vGroup.PUT("/a-items/:id", h.Update)
			vGroup.DELETE("/a-items/:id", h.Delete)
		}
	}
}

func (h *Handler) Run(port string) error {
	return h.config.Router.Run(port)
}
