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
	Update(ctx context.Context, id uuid.UUID, item domain.ItemB) error
	Remove(ctx context.Context, id uuid.UUID) error
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

func (r repository) GetAll(ctx context.Context) ([]*domain.ItemB, error) {
	cacheData, err := r.cache.Get("all-itemA")
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewArrayFromBytes(cacheData)
	}

	var itemArr []*domain.ItemB
	if err = r.database.Select(&itemArr); err != nil {
		return nil, err
	}

	if err = r.cache.Set("all-itemA", itemArr); err != nil {
		return nil, err
	}

	return itemArr, nil
}

func (r repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemB, error) {
	cacheData, err := r.cache.Get(id.String())
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewFromBytes(cacheData)
	}

	item := &domain.ItemB{ID: id}
	if err = r.database.Select(item); err != nil {
		return nil, err
	}

	if err = r.cache.Set(id.String(), item); err != nil {
		return nil, err
	}

	return item, nil
}

func (r repository) Insert(ctx context.Context, item domain.ItemB) (*domain.ItemB, error) {
	err := r.cache.Remove("all-itemA")
	if err != nil {
		return nil, err
	}

	input := &item
	if err = r.database.Create(input); err != nil {
		return nil, err
	}

	return input, nil
}

func (r repository) Update(ctx context.Context, id uuid.UUID, item domain.ItemB) error {
	err := r.cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.cache.Remove("all-itemA")
	if err != nil {
		return err
	}

	return r.database.Update(id, item)
}

func (r repository) Remove(ctx context.Context, id uuid.UUID) error {
	err := r.cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.cache.Remove("all-itemA")
	if err != nil {
		return err
	}

	return r.database.Delete(id, domain.ItemB{})
}
