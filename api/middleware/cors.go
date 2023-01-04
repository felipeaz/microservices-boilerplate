package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	ginCors "github.com/itsjamie/gin-cors"
)

type cors struct{}

func NewCorsMiddleware() Middleware {
	return &cors{}
}

func (m *cors) HandleFunc() gin.HandlerFunc {
	return ginCors.Middleware(ginCors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     false,
		ValidateHeaders: false,
	})
}
