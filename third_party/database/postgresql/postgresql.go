package postgresql

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/internal/storage"
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

func (p *postgresql) Create(ctx context.Context, obj interface{}) error {
	conn, db := p.connect()
	defer p.closeConnection(db)
	return conn.WithContext(ctx).Create(obj).Error
}

func (p *postgresql) Update(ctx context.Context, id uuid.UUID, obj interface{}) error {
	conn, db := p.connect()
	defer p.closeConnection(db)
	return conn.WithContext(ctx).Where("id = ?", id).Updates(obj).Error
}

func (p *postgresql) Set(ctx context.Context, obj interface{}, field string, value interface{}) error {
	conn, db := p.connect()
	defer p.closeConnection(db)
	return conn.WithContext(ctx).Model(obj).UpdateColumn(field, value).Error
}

func (p *postgresql) Select(ctx context.Context, obj interface{}) error {
	conn, db := p.connect()
	defer p.closeConnection(db)
	return conn.WithContext(ctx).Find(obj).Error
}

func (p *postgresql) Raw(ctx context.Context, query string, obj interface{}) error {
	conn, db := p.connect()
	defer p.closeConnection(db)
	return conn.WithContext(ctx).Raw(query).Scan(obj).Error
}

func (p *postgresql) Delete(ctx context.Context, id uuid.UUID, obj interface{}) error {
	conn, db := p.connect()
	defer p.closeConnection(db)
	return conn.WithContext(ctx).Where("id = ?", id).Delete(obj).Error
}

func (p *postgresql) getDBInfo() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC",
		p.host, p.port, p.username, p.password, p.dbName,
	)
}

func (p *postgresql) closeConnection(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Println("failed to close sql db connection", err)
	}
}
