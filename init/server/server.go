package server

import (
	"log"

	"microservices-boilerplate/api"
)

type Server interface {
	Run(addr ...string)
}

func New(api api.Api) Server {
	return &server{
		api: api,
	}
}

type server struct {
	api api.Api
}

func (s *server) Run(addr ...string) {
	err := s.api.GetRouter().Run(addr...)
	if err != nil {
		log.Fatal("failed to initialize server")
	}
}
