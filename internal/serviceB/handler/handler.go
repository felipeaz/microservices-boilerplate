package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	httpService "microservices-boilerplate/internal/http"
	"microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/internal/serviceB/service"
)

type Handler struct {
	service   service.Service
	httpError httpService.Error
}

func New(service service.Service) *Handler {
	return &Handler{
		service:   service,
		httpError: httpService.NewHttpError(),
	}
}

// Get godoc
// @Summary     Show all items
// @Description Return all stored items
// @Tags        itemB
// @Accept      json
// @Produce     json
// @Success     200 {array}  domain.ItemB
// @Failure     500   {object} error
// @Router      /b-items [get]
func (h *Handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	resp, err := h.service.GetAll(ctx)
	if err != nil {
		h.httpError.GetStatusCodeFromError(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Find godoc
// @Summary     Show an item
// @Description get item by ID
// @Tags        itemB
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Item ID"
// @Success     200   {object} domain.ItemB
// @Failure     400   {object} error
// @Failure     404   {object} error
// @Failure     500 {object} error
// @Router      /b-items/{id} [get]
func (h *Handler) Find(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	resp, err := h.service.GetOneByID(ctx, id)
	if err != nil {
		h.httpError.GetStatusCodeFromError(err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create godoc
// @Summary     Creates an item
// @Description creates an item with given data
// @Tags        itemB
// @Accept      json
// @Produce     json
// @Param       itemB body domain.ItemB true "Item Properties"
// @Success     200 {object} domain.ItemB
// @Failure     400 {object} error
// @Failure     404 {object} error
// @Failure     500 {object} error
// @Router      /b-items [post]
func (h *Handler) Create(c *gin.Context) {
	var input *domain.ItemB
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	obj, err := h.service.Create(ctx, input)
	if err != nil {
		h.httpError.GetStatusCodeFromError(err)
		return
	}

	c.JSON(http.StatusOK, obj)
}

// Update godoc
// @Summary     Updates an item
// @Description Updates an item with given ID
// @Tags        itemB
// @Accept      json
// @Produce     json
// @Param       id path string true "Item ID"
// @Param       itemB body domain.ItemB true "Item Properties"
// @Success     200
// @Failure     400 {object} error
// @Failure     404 {object} error
// @Failure     500 {object} error
// @Router      /b-items/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	var input *domain.ItemB
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	id := c.Param("id")
	if err = h.service.Update(ctx, id, input); err != nil {
		h.httpError.GetStatusCodeFromError(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Delete godoc
// @Summary     Deletes an item
// @Description Deletes an item with given ID
// @Tags        itemB
// @Accept      json
// @Produce     json
// @Param       string path     string true "Item ID"
// @Success     200    {object} domain.ItemB
// @Failure     400    {object} error
// @Failure     404    {object} error
// @Failure     500    {object} error
// @Router      /b-items/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := h.service.Delete(ctx, id); err != nil {
		h.httpError.GetStatusCodeFromError(err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
