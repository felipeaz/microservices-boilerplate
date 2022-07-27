package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	httpService "microservices-boilerplate/internal/http"
	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/serviceA/service"
)

type Handler struct {
	service   service.Service
	httpError httpService.Error
}

func New(service service.Service) Handler {
	return Handler{
		service:   service,
		httpError: httpService.NewHttpError(),
	}
}

func (h Handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	resp, err := h.service.GetAll(ctx)
	if err != nil {
		c.JSON(h.httpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h Handler) Find(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	resp, err := h.service.GetOneByID(ctx, id)
	if err != nil {
		c.JSON(h.httpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h Handler) Create(c *gin.Context) {
	var input *domain.ItemA
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	obj, err := h.service.Create(ctx, input)
	if err != nil {
		c.JSON(h.httpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusOK, obj)
}

func (h Handler) Update(c *gin.Context) {
	var input *domain.ItemA
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	id := c.Param("id")
	if err = h.service.Update(ctx, id, input); err != nil {
		c.JSON(h.httpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := h.service.Delete(ctx, id); err != nil {
		c.JSON(h.httpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
