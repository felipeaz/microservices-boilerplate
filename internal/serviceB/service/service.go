package service

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"app/internal/constants"
	"app/internal/logger"
	"app/internal/serviceB/domain"
	"app/internal/serviceB/repository"
	"app/internal/serviceB/service/metrics"
)

type Service interface {
	GetAll(ctx context.Context) ([]*domain.ItemB, error)
	GetOneByID(ctx context.Context, id string) (*domain.ItemB, error)
	Create(ctx context.Context, item *domain.ItemB) (*domain.ItemB, error)
	Update(ctx context.Context, id string, item *domain.ItemB) error
	Delete(ctx context.Context, id string) error
}

type DependenciesNode struct {
	Log        logger.Logger
	Repository repository.Repository
}

type service struct {
	deps    *DependenciesNode
	metrics *metrics.Metrics
}

func New(config *DependenciesNode) Service {
	return &service{
		deps:    config,
		metrics: metrics.Initialize(),
	}
}

func (s *service) GetAll(ctx context.Context) ([]*domain.ItemB, error) {
	resp, err := s.deps.Repository.GetAll(ctx)
	if err != nil {
		s.deps.Log.Error(FailedToGetAll, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemB, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.deps.Log.Error(FailedToParseUUID, err)
		return nil, constants.ErrCreatingUUIDFromString
	}

	resp, err := s.deps.Repository.GetByID(ctx, itemID)
	if err != nil {
		s.deps.Log.Error(FailedToGetByID, itemID, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemB) (*domain.ItemB, error) {
	resp, err := s.deps.Repository.Insert(ctx, item)
	if err != nil {
		s.deps.Log.Error(FailedToCreate, item, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemB) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.deps.Log.Error(FailedToParseUUID, err)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Update(ctx, itemID, item); err != nil {
		s.deps.Log.Error(FailedToUpdate, itemID, item, err)
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.deps.Log.Error(FailedToParseUUID, err)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Remove(ctx, itemID); err != nil {
		s.deps.Log.Error(FailedToDelete, itemID, err)
		return err
	}

	return nil
}
