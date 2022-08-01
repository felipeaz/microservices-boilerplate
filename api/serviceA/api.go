package serviceA

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	. "microservices-boilerplate/api"
	"microservices-boilerplate/internal/serviceA/handler"
)

type api struct {
	handler handler.Handler
	router  *gin.Engine
}

func NewApi(h handler.Handler, router *gin.Engine) Api {
	return &api{
		handler: h,
		router:  router,
	}
}

func (a *api) RegisterRoutes() {
	apiGroup := a.router.Group("/api")
	vGroup := apiGroup.Group("/v1")

	vGroup.GET("/item", a.handler.Get)
	vGroup.GET("/item/:id", a.handler.Find)
	vGroup.POST("/item", a.handler.Create)
	vGroup.PUT("/item/:id", a.handler.Update)
	vGroup.DELETE("/item/:id", a.handler.Delete)

	a.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func (a *api) Run(port string) error {
	return a.router.Run(port)
}
