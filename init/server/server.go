package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"microservices-boilerplate/api"
)

type Server interface {
	Run(port string)
}

func New(api api.Api) Server {
	return &server{
		api: api,
	}
}

type server struct {
	api api.Api
}

func (s *server) Run(port string) {
	router := gin.Default()

	s.api.RegisterRoutes(router)

	err := router.Run(port)
	if err != nil {
		log.Fatal("failed to initialize server")
	}
}
