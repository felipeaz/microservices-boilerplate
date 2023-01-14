package repository

import (
	"context"
	"time"

	uuid "github.com/satori/go.uuid"

	"app/internal/serviceA/domain"
	"app/internal/serviceA/repository/metrics"
	"app/internal/storage"
)

const (
	AllItemsKey = "all-items"

	cachedQueryMetric = "cached"
	dbQueryMetric     = "db"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*domain.ItemA, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error)
	Insert(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id uuid.UUID, item *domain.ItemA) error
	Remove(ctx context.Context, id uuid.UUID) error
}

type DependenciesNode struct {
	Database storage.Database
	Cache    storage.Cache
}

type repository struct {
	deps    *DependenciesNode
	metrics *metrics.Metrics
}

func New(deps *DependenciesNode) Repository {
	return &repository{
		deps:    deps,
		metrics: metrics.Initialize(),
	}
}

func (r *repository) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	startTime := time.Now()
	cacheData, err := r.deps.Cache.Get(AllItemsKey)
	if err != nil {
		return nil, err
	}

	if cacheData != nil {
		r.metrics.Latency.Observe(time.Since(startTime).Seconds(), cachedQueryMetric)
		return domain.NewArrayFromBytes(cacheData)
	}

	var itemArr []*domain.ItemA
	if err = r.deps.Database.Select(ctx, &itemArr); err != nil {
		return nil, err
	}

	if err = r.deps.Cache.Set(AllItemsKey, itemArr); err != nil {
		return nil, err
	}

	r.metrics.Latency.Observe(time.Since(startTime).Seconds(), dbQueryMetric)

	return itemArr, nil
}

func (r *repository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ItemA, error) {
	startTime := time.Now()
	cacheData, err := r.deps.Cache.Get(id.String())
	if err != nil {
		return nil, err
	}

	if cacheData != nil {
		r.metrics.Latency.Observe(time.Since(startTime).Seconds(), cachedQueryMetric)
		return domain.NewFromBytes(cacheData)
	}

	item := &domain.ItemA{ID: id}
	if err = r.deps.Database.Select(ctx, item); err != nil {
		return nil, err
	}

	if err = r.deps.Cache.Set(id.String(), item); err != nil {
		return nil, err
	}

	r.metrics.Latency.Observe(time.Since(startTime).Seconds(), dbQueryMetric)

	return item, nil
}

func (r *repository) Insert(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	startTime := time.Now()
	err := r.deps.Cache.Remove(AllItemsKey)
	if err != nil {
		return nil, err
	}

	if err = r.deps.Database.Create(ctx, item); err != nil {
		return nil, err
	}

	r.metrics.Latency.Observe(time.Since(startTime).Seconds(), dbQueryMetric)

	return item, nil
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, item *domain.ItemA) error {
	startTime := time.Now()
	err := r.deps.Cache.Remove(id.String())
	if err != nil {
		return err
	}

	err = r.deps.Cache.Remove(AllItemsKey)
	if err != nil {
		return err
	}

	r.metrics.Latency.Observe(time.Since(startTime).Seconds(), dbQueryMetric)

	return r.deps.Database.Update(ctx, id, item)
}

func (r *repository) Remove(ctx context.Context, id uuid.UUID) error {
	startTime := time.Now()
	err := r.deps.Cache.Remove(id.String())
	if err != nil {
		return err
	}
	err = r.deps.Cache.Remove(AllItemsKey)
	if err != nil {
		return err
	}

	r.metrics.Latency.Observe(time.Since(startTime).Seconds(), dbQueryMetric)

	return r.deps.Database.Delete(ctx, id, domain.ItemA{})
}
