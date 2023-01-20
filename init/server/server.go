package server

import (
	"github.com/gin-gonic/gin"
)

type Server interface {
	Run(addr ...string) error
}

func New(httpEngine *gin.Engine) Server {
	return &server{
		httpEngine: httpEngine,
	}
}

type server struct {
	httpEngine *gin.Engine
}

func (s *server) Run(addr ...string) error {
	return s.httpEngine.Run(addr...)
}
