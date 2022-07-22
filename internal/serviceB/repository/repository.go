package repository

import (
	"context"
	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/internal/storage"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.ItemB, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemB, error)
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

func (r repository) GetAll(ctx context.Context) ([]*domain.ItemB, error) {
	cacheData, err := r.cache.Get("all-itemB")
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewArrayFromBytes(cacheData)
	}

	//TODO implement me
	panic("implement me")
}

func (r repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemB, error) {
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

func (r repository) Insert(ctx context.Context, item domain.ItemB) (*domain.ItemB, error) {
	err := r.cache.Remove("all-itemB")
	if err != nil {
		return nil, err
	}

	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, id uuid.UUID, item domain.ItemB) (*domain.ItemB, error) {
	err := r.cache.Remove(id.String())
	if err != nil {
		return nil, err
	}
	err = r.cache.Remove("all-itemB")
	if err != nil {
		return nil, err
	}

	//TODO implement me
	panic("implement me")
}

func (r repository) Remove(ctx context.Context, id uuid.UUID) (*domain.ItemB, error) {
	err := r.cache.Remove(id.String())
	if err != nil {
		return nil, err
	}
	err = r.cache.Remove("all-itemB")
	if err != nil {
		return nil, err
	}

	//TODO implement me
	panic("implement me")
}
