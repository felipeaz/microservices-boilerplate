package config

import (
	"app/build/env"
	"app/build/flags"
	"app/build/router"
	"app/internal/logger"
	"app/internal/storage"
	"app/third_party/cache/redis"
	"app/third_party/database/postgresql"
	"github.com/gin-gonic/gin"
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
