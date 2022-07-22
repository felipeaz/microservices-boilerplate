package repository

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/storage"
)

type Repository interface {
	Query(ctx context.Context, raw string) (*domain.ItemA, error)
	Insert(ctx context.Context, item domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id uuid.UUID, item domain.ItemA) (*domain.ItemA, error)
	Remove(ctx context.Context, id uuid.UUID) (*domain.ItemA, error)
}

func New(cache storage.Cache) Repository {
	return repository{
		cache: cache,
	}
}

type repository struct {
	cache storage.Cache
}

func (r repository) Query(ctx context.Context, raw string) (*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Insert(ctx context.Context, item domain.ItemA) (*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, id uuid.UUID, item domain.ItemA) (*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Remove(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}
