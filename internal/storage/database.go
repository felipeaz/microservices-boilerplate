package storage

import (
	"context"

	uuid "github.com/satori/go.uuid"
)

type Database interface {
	Create(ctx context.Context, obj interface{}) error
	Update(ctx context.Context, id uuid.UUID, obj interface{}) error
	Set(ctx context.Context, obj interface{}, field string, value interface{}) error
	Select(ctx context.Context, obj interface{}) error
	Raw(ctx context.Context, query string, obj interface{}) error
	Delete(ctx context.Context, id uuid.UUID, obj interface{}) error
}
