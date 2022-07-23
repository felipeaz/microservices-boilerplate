package env

import (
	"log"
	"os"
)

const (
	dbHostEnv     = "DB_HOST"
	dbUsernameEnv = "DB_USERNAME"
	dbPasswordEnv = "DB_PASSWORD"
	cacheHostEnv  = "CACHE_HOST"
	hostEnv       = "SERVER_HOST"

	missingEnvErr = "missing env: %s"
)

type Env struct {
	DBHost     string
	DBUsername string
	DBPassword string
	CacheHost  string
	Host       string
}

func Build() Env {
	var env Env
	var ok bool
	env.DBHost, ok = os.LookupEnv(dbHostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbHostEnv)
	}
	env.DBUsername, ok = os.LookupEnv(dbUsernameEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbUsernameEnv)
	}
	env.DBPassword, ok = os.LookupEnv(dbPasswordEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbPasswordEnv)
	}
	env.CacheHost, ok = os.LookupEnv(cacheHostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, cacheHostEnv)
	}
	env.Host, ok = os.LookupEnv(hostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, hostEnv)
	}
	return env
}
