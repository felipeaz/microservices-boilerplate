package postgresql

import (
	"database/sql"
	"fmt"
	uuid "github.com/satori/go.uuid"
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

func (p *postgresql) connect() (*gorm.DB, *sql.DB) {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.getDBInfo(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to postgresql: %v\n", err)
	}
	db, err := conn.DB()
	if err != nil {
		log.Fatalf("failed to retrieve db instance: %v\n", err)
	}
	return conn, db
}

func (p *postgresql) Create(obj interface{}) error {
	conn, db := p.connect()
	defer db.Close()
	return conn.Create(obj).Error
}

func (p *postgresql) Update(id uuid.UUID, obj interface{}) error {
	conn, db := p.connect()
	defer db.Close()
	return conn.Where("id = ?", id).Updates(obj).Error
}

func (p *postgresql) Set(obj interface{}, field string, value interface{}) error {
	conn, db := p.connect()
	defer db.Close()
	return conn.Model(obj).UpdateColumn(field, value).Error
}

func (p *postgresql) Select(obj interface{}) error {
	conn, db := p.connect()
	defer db.Close()
	return conn.Find(obj).Error
}

func (p *postgresql) Raw(query string, obj interface{}) error {
	conn, db := p.connect()
	defer db.Close()
	return conn.Raw(query).Scan(obj).Error
}

func (p *postgresql) Delete(id uuid.UUID, obj interface{}) error {
	conn, db := p.connect()
	defer db.Close()
	return conn.Where("id = ?", id).Delete(obj).Error
}

func (p *postgresql) getDBInfo() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		p.host, p.port, p.username, p.password, p.dbName,
	)
}
