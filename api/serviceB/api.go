package serviceB

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	. "microservices-boilerplate/api"
	"microservices-boilerplate/internal/serviceB/handler"
)

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
	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiGroup := a.router.Group("/api")
	{
		vGroup := apiGroup.Group("/v1")
		{
			vGroup.GET("/b-items", a.handler.Get)
			vGroup.GET("/b-items/:id", a.handler.Find)
			vGroup.POST("/b-items", a.handler.Create)
			vGroup.PUT("/b-items/:id", a.handler.Update)
			vGroup.DELETE("/b-items/:id", a.handler.Delete)
		}
	}
}

func (a *api) Run(port string) error {
	return a.router.Run(port)
}
