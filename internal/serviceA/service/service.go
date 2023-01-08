package service

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"app/internal/constants"
	"app/internal/logger"
	"app/internal/metrics"
	"app/internal/serviceA/domain"
	"app/internal/serviceA/repository"
)

type Service interface {
	GetAll(ctx context.Context) ([]*domain.ItemA, error)
	GetOneByID(ctx context.Context, id string) (*domain.ItemA, error)
	Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id string, item *domain.ItemA) error
	Delete(ctx context.Context, id string) error
}

type DependenciesNode struct {
	Repository repository.Repository
	Log        logger.Logger
}

type service struct {
	deps    *DependenciesNode
	metrics metrics.MetricCollector
}

func New(deps *DependenciesNode) Service {
	return &service{
		deps:    deps,
		metrics: metrics.NewMetricCollector(),
	}
}

func (s *service) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	resp, err := s.deps.Repository.GetAll(ctx)
	if err != nil {
		s.deps.Log.Error("failed to get all item a", err)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemA, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.deps.Log.Error("failed to parse id to UUID", err)
		return nil, constants.ErrCreatingUUIDFromString
	}

	resp, err := s.deps.Repository.GetByID(ctx, itemID)
	if err != nil {
		s.deps.Log.Error("failed to get item with id", itemID, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	resp, err := s.deps.Repository.Insert(ctx, item)
	if err != nil {
		s.deps.Log.Error("failed to create item", item, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemA) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.deps.Log.Error("failed to parse id to UUID", err)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Update(ctx, itemID, item); err != nil {
		s.deps.Log.Error("failed to update item", itemID, item, err)
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.deps.Log.Error("failed to parse id to UUID", err)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Remove(ctx, itemID); err != nil {
		s.deps.Log.Error("failed to delete item", itemID, err)
		return err
	}

	return nil
}
