package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"microservices-boilerplate/api"
	"microservices-boilerplate/api/middleware"
)

type Server interface {
	Run(port string)
}

func New(api api.Api, mw middleware.Middleware) Server {
	return &server{
		api:        api,
		middleware: mw,
	}
}

type server struct {
	api        api.Api
	middleware middleware.Middleware
}

func (s *server) Run(port string) {
	router := gin.Default()
	router.Use(s.middleware.Cors())

	s.api.RegisterRoutes(router)

	err := router.Run(port)
	if err != nil {
		log.Fatal("failed to initialize server")
	}
}
