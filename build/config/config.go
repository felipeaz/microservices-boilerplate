package config

import (
	"github.com/gin-gonic/gin"
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/build/router"
	"microservices-boilerplate/internal/logger"
	"microservices-boilerplate/internal/storage"
	"microservices-boilerplate/third_party/cache/redis"
	"microservices-boilerplate/third_party/database/postgresql"
)

type Config struct {
	Port     string
	Database storage.Database
	Cache    storage.Cache
	Logger   logger.Logger
	Router   *gin.Engine
}

func Build(env env.Env, flags flags.Flags) Config {
	return Config{
		Port:     env.Port,
		Database: postgresql.New(env.DBHost, env.DBPort, env.DBUsername, env.DBPassword, env.DBName),
		Cache:    redis.New(env.CacheHost, env.CachePort),
		Logger:   logger.NewLogger(*flags.Debug),
		Router:   router.New(),
	}
}
