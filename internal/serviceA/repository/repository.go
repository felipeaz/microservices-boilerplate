package repository

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/storage"
)

const (
	AllItemsKey = "all-items"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.ItemA, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error)
	Insert(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id uuid.UUID, item *domain.ItemA) error
	Remove(ctx context.Context, id uuid.UUID) error
}

func New(db storage.Database, cache storage.Cache) Repository {
	return &repository{
		database: db,
		cache:    cache,
	}
}

type repository struct {
	database storage.Database
	cache    storage.Cache
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	cacheData, err := r.cache.Get(AllItemsKey)
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewArrayFromBytes(cacheData)
	}

	var itemArr []*domain.ItemA
	if err = r.database.Select(ctx, &itemArr); err != nil {
		return nil, err
	}

	if err = r.cache.Set(AllItemsKey, itemArr); err != nil {
		return nil, err
	}

	return itemArr, nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	cacheData, err := r.cache.Get(id.String())
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewFromBytes(cacheData)
	}

	item := &domain.ItemA{ID: id}
	if err = r.database.Select(ctx, item); err != nil {
		return nil, err
	}

	if err = r.cache.Set(id.String(), item); err != nil {
		return nil, err
	}

	return item, nil
}

func (r *repository) Insert(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	err := r.cache.Remove(AllItemsKey)
	if err != nil {
		return nil, err
	}

	if err = r.database.Create(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, item *domain.ItemA) error {
	err := r.cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.cache.Remove(AllItemsKey)
	if err != nil {
		return err
	}

	return r.database.Update(ctx, id, item)
}

func (r *repository) Remove(ctx context.Context, id uuid.UUID) error {
	err := r.cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.cache.Remove(AllItemsKey)
	if err != nil {
		return err
	}

	return r.database.Delete(ctx, id, domain.ItemA{})
}
