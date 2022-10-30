package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"microservices-boilerplate/api"
	httpService "microservices-boilerplate/internal/http"
	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/serviceA/service"
)

type DependenciesNode struct {
	Service   service.Service
	HttpError httpService.Error
	Router    *gin.Engine
}

type Handler struct {
	deps *DependenciesNode
}

func New(deps *DependenciesNode) *Handler {
	handler := &Handler{
		deps: deps,
	}
	handler.RegisterRoutes()
	return handler
}

func (h *Handler) RegisterRoutes() {
	h.deps.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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

func (h *Handler) GetRouter() api.Router {
	return h.deps.Router
}

// Get godoc
// @Summary     Show all items
// @Description Return all stored items
// @Tags        itemA
// @Accept      json
// @Produce     json
// @Success     200 {array}  domain.ItemA
// @Failure     500   {object} error
// @Router      /a-items [get]
func (h *Handler) Get(c *gin.Context) {
	ctx := c.Request.Context()
	resp, err := h.deps.Service.GetAll(ctx)
	if err != nil {
		c.JSON(h.deps.HttpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Find godoc
// @Summary     Show an item
// @Description get item by ID
// @Tags        itemA
// @Accept      json
// @Produce     json
// @Param       id  path     string true "Item ID"
// @Success     200   {object} domain.ItemA
// @Failure     400   {object} error
// @Failure     404 {object} error
// @Failure     500 {object} error
// @Router      /a-items/{id} [get]
func (h *Handler) Find(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	resp, err := h.deps.Service.GetOneByID(ctx, id)
	if err != nil {
		c.JSON(h.deps.HttpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusOK, resp)
}

// Create godoc
// @Summary     Creates an item
// @Description creates an item with given data
// @Tags        itemA
// @Accept      json
// @Produce     json
// @Param       itemA body     domain.ItemA true "Item Properties"
// @Success     200 {object} domain.ItemA
// @Failure     400 {object} error
// @Failure     500 {object} error
// @Router      /a-items [post]
func (h *Handler) Create(c *gin.Context) {
	var input *domain.ItemA
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	obj, err := h.deps.Service.Create(ctx, input)
	if err != nil {
		c.JSON(h.deps.HttpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusOK, obj)
}

// Update godoc
// @Summary     Updates an item
// @Description Updates an item with given ID
// @Tags        itemA
// @Accept      json
// @Produce     json
// @Param       id    path string       true "Item ID"
// @Param       itemA body domain.ItemA true "Item Properties"
// @Success     200
// @Failure     400 {object} error
// @Failure     404 {object} error
// @Failure     500 {object} error
// @Router      /a-items/{id} [put]
func (h *Handler) Update(c *gin.Context) {
	var input *domain.ItemA
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx := c.Request.Context()
	id := c.Param("id")
	if err = h.deps.Service.Update(ctx, id, input); err != nil {
		c.JSON(h.deps.HttpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// Delete godoc
// @Summary     Deletes an item
// @Description Deletes an item with given ID
// @Tags        itemA
// @Accept      json
// @Produce     json
// @Param       string path     string true "Item ID"
// @Success     200    {object} domain.ItemA
// @Failure     400    {object} error
// @Failure     404    {object} error
// @Failure     500    {object} error
// @Router      /a-items/{id} [delete]
func (h *Handler) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if err := h.deps.Service.Delete(ctx, id); err != nil {
		c.JSON(h.deps.HttpError.GetStatusCodeFromError(err), err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
