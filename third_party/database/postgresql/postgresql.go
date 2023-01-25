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

const (
	dbInfo = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC"

	failedToConnectToPostgresql = "failed to connect to postgresql: %v\n"
	failedToRetrieveDBInstance  = "failed to retrieve db instance: %v\n"
)

type postgresql struct {
	host     string
	port     string
	username string
	password string
	dbName   string
	conn     *gorm.DB
	db       *sql.DB
}

func New(host, port, user, pass, dbName string) storage.Database {
	pgsql := &postgresql{
		host:     host,
		port:     port,
		username: user,
		password: pass,
		dbName:   dbName,
	}
	pgsql.connect()

	return pgsql
}

func (p *postgresql) connect() {
	if p.conn != nil {
		return
	}

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.getDBInfo(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf(failedToConnectToPostgresql, err)
	}

	db, err := conn.DB()
	if err != nil {
		log.Fatalf(failedToRetrieveDBInstance, err)
	}

	p.conn = conn
	p.db = db
}

func (p *postgresql) Create(ctx context.Context, obj interface{}) error {
	return p.conn.WithContext(ctx).Create(obj).Error
}

func (p *postgresql) Update(ctx context.Context, id uuid.UUID, obj interface{}) error {
	return p.conn.WithContext(ctx).Where("id = ?", id).Updates(obj).Error
}

func (p *postgresql) Set(ctx context.Context, obj interface{}, field string, value interface{}) error {
	return p.conn.WithContext(ctx).Model(obj).UpdateColumn(field, value).Error
}

func (p *postgresql) Select(ctx context.Context, obj interface{}) error {
	return p.conn.WithContext(ctx).Find(obj).Error
}

func (p *postgresql) Raw(ctx context.Context, query string, obj interface{}) error {
	return p.conn.WithContext(ctx).Raw(query).Scan(obj).Error
}

func (p *postgresql) Delete(ctx context.Context, id uuid.UUID, obj interface{}) error {
	return p.conn.WithContext(ctx).Where("id = ?", id).Delete(obj).Error
}

func (p *postgresql) getDBInfo() string {
	return fmt.Sprintf(dbInfo, p.host, p.port, p.username, p.password, p.dbName)
}
