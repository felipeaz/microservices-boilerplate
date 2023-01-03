package handler

import (
	"microservices-boilerplate/api"
)

func (h *Handler) GetRouter() api.Router {
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
			vGroup.GET("/a-items", h.Get)
			vGroup.GET("/a-items/:id", h.Find)
			vGroup.POST("/a-items", h.Create)
			vGroup.PUT("/a-items/:id", h.Update)
			vGroup.DELETE("/a-items/:id", h.Delete)
		}
	}
}
