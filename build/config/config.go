package config

import (
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/storage"
	"microservices-boilerplate/internal/storage/postgresql"
	"microservices-boilerplate/internal/storage/redis"
)

type Config struct {
	Port     string
	Database storage.Database
	Cache    storage.Cache
	Log      pkg.Logger
}

func Build(env env.Env, flags flags.Flags) Config {
	return Config{
		Port:     env.Port,
		Database: postgresql.New(env.DBHost, env.DBPort, env.DBUsername, env.DBPassword, env.DBName),
		Cache:    redis.New(env.CacheHost, env.CachePort),
		Log:      pkg.NewLogger(*flags.Debug),
	}
}
