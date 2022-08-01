package serviceB

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	. "microservices-boilerplate/api"
	"microservices-boilerplate/internal/serviceB/handler"
)

// @title       Service B Swagger Example API
// @version     1.0
// @description This is a sample server.

// @host     localhost:8086
// @BasePath /api/v1

type api struct {
	handler *handler.Handler
	router  *gin.Engine
}

func NewApi(h *handler.Handler, router *gin.Engine) Api {
	return &api{
		handler: h,
		router:  router,
	}
}

func (a *api) RegisterRoutes() {
	apiGroup := a.router.Group("/api")
	vGroup := apiGroup.Group("/v1")

	vGroup.GET("/items", a.handler.Get)
	vGroup.GET("/items/:id", a.handler.Find)
	vGroup.POST("/items", a.handler.Create)
	vGroup.PUT("/items/:id", a.handler.Update)
	vGroup.DELETE("/items/:id", a.handler.Delete)

	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (a *api) Run(port string) error {
	return a.router.Run(port)
}
