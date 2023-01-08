package env

import (
	"log"
	"os"
)

const (
	dbHostEnv     = "DB_HOST"
	dbPortEnv     = "DB_PORT"
	dbUsernameEnv = "DB_USERNAME"
	dbPasswordEnv = "DB_PASSWORD"
	dbNameEnv     = "DB_NAME"
	cacheHostEnv  = "CACHE_HOST"
	cachePortEnv  = "CACHE_PORT"
	hostEnv       = "SERVER_HOST"
	portEnv       = "SERVER_PORT"

	missingEnvErr = "missing env: %s"
)

type Env struct {
	DBEnv      DBEnv
	CacheEnv   CacheEnv
	ServiceEnv ServiceEnv
}

func Build() Env {
	var env Env
	var ok bool
	env.DBEnv.Server.Host, ok = os.LookupEnv(dbHostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbHostEnv)
	}
	env.DBEnv.Server.Port, ok = os.LookupEnv(dbPortEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbPortEnv)
	}
	env.DBEnv.Credentials.Username, ok = os.LookupEnv(dbUsernameEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbUsernameEnv)
	}
	env.DBEnv.Credentials.Password, ok = os.LookupEnv(dbPasswordEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbPasswordEnv)
	}
	env.DBEnv.DatabaseName, ok = os.LookupEnv(dbNameEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbNameEnv)
	}
	env.CacheEnv.Server.Host, ok = os.LookupEnv(cacheHostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, cacheHostEnv)
	}
	env.CacheEnv.Server.Port, ok = os.LookupEnv(cachePortEnv)
	if !ok {
		log.Fatalf(missingEnvErr, cachePortEnv)
	}
	env.ServiceEnv.Server.Host, ok = os.LookupEnv(hostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, hostEnv)
	}
	env.ServiceEnv.Server.Port, ok = os.LookupEnv(portEnv)
	if !ok {
		log.Fatalf(missingEnvErr, portEnv)
	}
	return env
}
