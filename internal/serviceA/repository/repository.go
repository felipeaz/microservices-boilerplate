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

func New(config *Config) Repository {
	return &repository{
		config: config,
	}
}

type Config struct {
	Database storage.Database
	Cache    storage.Cache
}

type repository struct {
	config *Config
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	cacheData, err := r.config.Cache.Get(AllItemsKey)
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewArrayFromBytes(cacheData)
	}

	var itemArr []*domain.ItemA
	if err = r.config.Database.Select(ctx, &itemArr); err != nil {
		return nil, err
	}

	if err = r.config.Cache.Set(AllItemsKey, itemArr); err != nil {
		return nil, err
	}

	return itemArr, nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	cacheData, err := r.config.Cache.Get(id.String())
	if err != nil {
		return nil, err
	}
	if cacheData != nil {
		return domain.NewFromBytes(cacheData)
	}

	item := &domain.ItemA{ID: id}
	if err = r.config.Database.Select(ctx, item); err != nil {
		return nil, err
	}

	if err = r.config.Cache.Set(id.String(), item); err != nil {
		return nil, err
	}

	return item, nil
}

func (r *repository) Insert(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	err := r.config.Cache.Remove(AllItemsKey)
	if err != nil {
		return nil, err
	}

	if err = r.config.Database.Create(ctx, item); err != nil {
		return nil, err
	}

	return item, nil
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, item *domain.ItemA) error {
	err := r.config.Cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.config.Cache.Remove(AllItemsKey)
	if err != nil {
		return err
	}

	return r.config.Database.Update(ctx, id, item)
}

func (r *repository) Remove(ctx context.Context, id uuid.UUID) error {
	err := r.config.Cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.config.Cache.Remove(AllItemsKey)
	if err != nil {
		return err
	}

	return r.config.Database.Delete(ctx, id, domain.ItemA{})
}
