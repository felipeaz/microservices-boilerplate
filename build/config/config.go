package config

import (
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/internal/pkg/log"
	"microservices-boilerplate/internal/storage"
	"microservices-boilerplate/third_party/postgresql"
	"microservices-boilerplate/third_party/redis"
	"time"
)

type Config struct {
	Port     string
	Database storage.Database
	Cache    storage.Cache
	Logger   log.Logger
}

func Build(env env.Env, flags flags.Flags) Config {
	return Config{
		Port:     env.Port,
		Database: postgresql.New(env.DBHost, env.DBPort, env.DBUsername, env.DBPassword, env.DBName),
		Cache:    redis.New(env.CacheHost, env.CachePort),
		Logger:   log.NewLogger(log.NewLogFile(time.Now(), log.GetLogPath()), *flags.Debug),
	}
}
