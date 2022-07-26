package postgresql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"microservices-boilerplate/internal/storage"
)

type postgresql struct {
	host     string
	port     string
	username string
	password string
	dbName   string
}

func New(host, port, user, pass, dbName string) storage.Database {
	return &postgresql{
		host:     host,
		port:     port,
		username: user,
		password: pass,
		dbName:   dbName,
	}
}

func (p *postgresql) connect() *gorm.DB {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.getDBInfo(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to postgresql: %v\n", err)
	}
	return conn
}

func (p *postgresql) getDBInfo() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		p.host, p.port, p.username, p.password, p.dbName,
	)
}
