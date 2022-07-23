package server

import (
	"log"

	"github.com/gin-gonic/gin"

	"microservices-boilerplate/api"
)

type Server interface {
	Run(addr string)
}

func New(api api.Api) Server {
	return &server{
		api: api,
	}
}

type server struct {
	api api.Api
}

func (s *server) Run(addr string) {
	router := gin.Default()

	s.api.RegisterRoutes(router)

	err := router.Run(addr)
	if err != nil {
		log.Fatal("failed to initialize server")
	}
}
