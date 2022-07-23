package config

import (
	"microservices-boilerplate/build/env"
	"microservices-boilerplate/build/flags"
)

type Config struct {
	Env   env.Env
	Flags flags.Flags
}

func Build() Config {
	return Config{
		Env:   env.Build(),
		Flags: flags.Build(),
	}
}
