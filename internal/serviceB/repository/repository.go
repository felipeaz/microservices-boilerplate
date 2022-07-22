package repository

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/internal/storage"
)

type Repository interface {
	Query(ctx context.Context, raw string) (*domain.ItemB, error)
	Insert(ctx context.Context, item domain.ItemB) (*domain.ItemB, error)
	Update(ctx context.Context, id uuid.UUID, item domain.ItemB) (*domain.ItemB, error)
	Remove(ctx context.Context, id uuid.UUID) (*domain.ItemB, error)
}

func New(cache storage.Cache) Repository {
	return repository{
		cache: cache,
	}
}

type repository struct {
	cache storage.Cache
}

func (r repository) Query(ctx context.Context, raw string) (*domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Insert(ctx context.Context, item domain.ItemB) (*domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, id uuid.UUID, item domain.ItemB) (*domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Remove(ctx context.Context, id uuid.UUID) (*domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}
