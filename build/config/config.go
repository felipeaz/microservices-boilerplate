package config

import (
	"app/build/env"
	"app/build/flags"
	"app/build/router"
	"app/infra/cache/redis"
	"app/infra/database/postgresql"
	"app/internal/logger"
	"app/internal/storage"
	"github.com/gin-gonic/gin"
)

type BuildArgs struct {
	Env    env.Env
	Flags  flags.Flags
	Router *gin.Engine
}

type Config struct {
	ServicePort string
	Database    storage.Database
	Cache       storage.Cache
	Logger      logger.Logger
	Router      *gin.Engine
}

func Build(args BuildArgs) Config {
	return Config{
		ServicePort: args.Env.ServiceEnv.Server.Port,
		Database: postgresql.New(
			args.Env.DBEnv.Server.Host,
			args.Env.DBEnv.Server.Port,
			args.Env.DBEnv.Credentials.Username,
			args.Env.DBEnv.Credentials.Password,
			args.Env.DBEnv.DatabaseName,
		),
		Cache: redis.New(
			args.Env.CacheEnv.Server.Host,
			args.Env.CacheEnv.Server.Port,
		),
		Logger: logger.NewLogger(
			*args.Flags.Debug,
		),
		Router: router.New(
			args.Router,
		),
	}
}
