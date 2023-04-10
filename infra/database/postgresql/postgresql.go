package postgresql

import (
	"app/infra/database/postgresql/executor"
	"context"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"app/internal/storage"
)

const (
	dbInfo = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC"

	failedToConnectToPostgresql = "failed to connect to postgresql: %v\n"
	unmappedExecutorErr         = "executor type %v is not mapped"
)

type postgresql struct {
	host     string
	port     string
	username string
	password string
	dbName   string
}

func New(host, port, user, pass, dbName string) storage.Database {
	pgsql := &postgresql{
		host:     host,
		port:     port,
		username: user,
		password: pass,
		dbName:   dbName,
	}

	return pgsql
}

func (p *postgresql) connect() (*gorm.DB, error) {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  p.getDBInfo(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return nil, connectionError(err)
	}
	return conn, nil
}

func (p *postgresql) Create(ctx context.Context, obj interface{}) error {
	return p.Exec(ctx, executor.ExecArgs{
		ExecutorType: executor.CreateType,
		Object:       obj,
	})
}

func (p *postgresql) Update(ctx context.Context, id uuid.UUID, obj interface{}) error {
	return p.Exec(ctx, executor.ExecArgs{
		ExecutorType: executor.UpdateType,
		ID:           id,
		Object:       obj,
	})
}

func (p *postgresql) Set(ctx context.Context, obj interface{}, field string, value interface{}) error {
	return p.Exec(ctx, executor.ExecArgs{
		ExecutorType: executor.SetType,
		Object:       obj,
		SetColumnArgs: executor.SetColumnArgs{
			Field: field,
			Value: value,
		},
	})
}

func (p *postgresql) Select(ctx context.Context, obj interface{}) error {
	return p.Exec(ctx, executor.ExecArgs{
		ExecutorType: executor.SelectType,
		Object:       obj,
	})
}

func (p *postgresql) Raw(ctx context.Context, query string, obj interface{}) error {
	return p.Exec(ctx, executor.ExecArgs{
		ExecutorType: executor.RawType,
		Object:       obj,
		QueryArgs: executor.QueryArgs{
			QueryString: query,
		},
	})
}

func (p *postgresql) Delete(ctx context.Context, id uuid.UUID, obj interface{}) error {
	return p.Exec(ctx, executor.ExecArgs{
		ExecutorType: executor.DeleteType,
		ID:           id,
		Object:       obj,
	})
}

func (p *postgresql) Exec(ctx context.Context, args executor.ExecArgs) error {
	conn, err := p.connect()
	if err != nil {
		return err
	}

	db, err := conn.DB()
	if err != nil {
		return err
	}

	dbExecutor := executor.NewExecutor(args.ExecutorType)
	if dbExecutor == nil {
		return fmt.Errorf(unmappedExecutorErr, args.ExecutorType)
	}

	err = dbExecutor.Exec(ctx, conn, args)
	if err != nil {
		return err
	}

	return db.Close()
}

func (p *postgresql) getDBInfo() string {
	return fmt.Sprintf(dbInfo, p.host, p.port, p.username, p.password, p.dbName)
}

func connectionError(err error) error {
	return fmt.Errorf(failedToConnectToPostgresql, err)
}
