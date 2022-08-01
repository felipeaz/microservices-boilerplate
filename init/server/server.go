package server

import (
	"log"

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
	s.api.RegisterRoutes()
	err := s.api.Run(port)
	if err != nil {
		log.Fatal("failed to initialize server")
	}
}
