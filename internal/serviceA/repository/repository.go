package repository

import (
	"context"
	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/storage"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.ItemA, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error)
	Insert(ctx context.Context, item domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id uuid.UUID, item domain.ItemA) (*domain.ItemA, error)
	Remove(ctx context.Context, id uuid.UUID) (*domain.ItemA, error)
}

func New(db storage.Database, cache storage.Cache) Repository {
	return repository{
		database: db,
		cache:    cache,
	}
}

type repository struct {
	database storage.Database
	cache    storage.Cache
}

func (r repository) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	cacheData, err := r.cache.Get("all-itemA")
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewArrayFromBytes(cacheData)
	}

	//TODO implement me
	panic("implement me")
}

func (r repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	cacheData, err := r.cache.Get(id.String())
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewFromBytes(cacheData)
	}

	//TODO implement me
	panic("implement me")
}

func (r repository) Insert(ctx context.Context, item domain.ItemA) (*domain.ItemA, error) {
	r.cache.Remove("all-itemA")

	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, id uuid.UUID, item domain.ItemA) (*domain.ItemA, error) {
	err := r.cache.Remove(id.String())
	if err != nil {
		return nil, err
	}
	err = r.cache.Remove("all-itemA")
	if err != nil {
		return nil, err
	}

	//TODO implement me
	panic("implement me")
}

func (r repository) Remove(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	err := r.cache.Remove(id.String())
	if err != nil {
		return nil, err
	}
	err = r.cache.Remove("all-itemA")
	if err != nil {
		return nil, err
	}

	//TODO implement me
	panic("implement me")
}
