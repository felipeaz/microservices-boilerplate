package handler

import (
	"github.com/gin-gonic/gin"

	"microservices-boilerplate/internal/serviceB/service"
)

type Handler struct {
	service service.Service
}

func New(service service.Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) Get(c *gin.Context) {

}

func (h Handler) Find(c *gin.Context) {

}

func (h Handler) Create(c *gin.Context) {

}

func (h Handler) Update(c *gin.Context) {

}

func (h Handler) Delete(c *gin.Context) {

}
