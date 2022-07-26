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
	cacheHostEnv  = "CACHE_HOST"
	cachePortEnv  = "CACHE_PORT"
	hostEnv       = "SERVER_HOST"
	portEnv       = "SERVER_PORT"

	missingEnvErr = "missing env: %s"
)

type Env struct {
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	CacheHost  string
	CachePort  string
	Host       string
	Port       string
}

func Build() Env {
	var env Env
	var ok bool
	env.DBHost, ok = os.LookupEnv(dbHostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbHostEnv)
	}
	env.DBPort, ok = os.LookupEnv(dbPortEnv)
	if !ok {
		log.Fatalf(missingEnvErr, dbPortEnv)
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
	env.CachePort, ok = os.LookupEnv(cachePortEnv)
	if !ok {
		log.Fatalf(missingEnvErr, cachePortEnv)
	}
	env.Host, ok = os.LookupEnv(hostEnv)
	if !ok {
		log.Fatalf(missingEnvErr, hostEnv)
	}
	env.Port, ok = os.LookupEnv(portEnv)
	if !ok {
		log.Fatalf(missingEnvErr, portEnv)
	}
	return env
}
