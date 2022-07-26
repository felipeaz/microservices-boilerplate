package postgresql

import "microservices-boilerplate/internal/storage"

type postgresql struct {
	addr     string
	port     string
	username string
	password string
}

func New(host, port, user, pass string) storage.Database {
	return postgresql{
		addr:     host,
		port:     port,
		username: user,
		password: pass,
	}
}
