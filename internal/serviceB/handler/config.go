package handler

import (
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetRouter() *gin.Engine {
	return h.deps.Router
}

func (h *Handler) RegisterRoutes() {
	h.registerApi()
}

func (h *Handler) registerApi() {
	apiGroup := h.deps.Router.Group("/api")
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
