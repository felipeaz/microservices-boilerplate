package service

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"app/internal/constants"
	"app/internal/logger"
	"app/internal/serviceA/domain"
	"app/internal/serviceA/repository"
	"app/internal/serviceA/service/metrics"
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
	metrics *metrics.Metrics
}

func New(deps *DependenciesNode) Service {
	return &service{
		deps:    deps,
		metrics: metrics.Initialize(),
	}
}

func (s *service) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	resp, err := s.deps.Repository.GetAll(ctx)
	if err != nil {
		s.handleError(err, FailedToGetAll)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemA, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(err, FailedToParseUUID, id)
		return nil, constants.ErrCreatingUUIDFromString
	}

	resp, err := s.deps.Repository.GetByID(ctx, itemID)
	if err != nil {
		s.handleError(err, FailedToGetByID, itemID)
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	resp, err := s.deps.Repository.Insert(ctx, item)
	if err != nil {
		s.handleError(err, FailedToCreate, item)
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemA) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(err, FailedToParseUUID, id)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Update(ctx, itemID, item); err != nil {
		s.handleError(err, FailedToUpdate, itemID, item)
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(err, FailedToParseUUID, id)
		return constants.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Remove(ctx, itemID); err != nil {
		s.handleError(err, FailedToDelete, itemID)
		return err
	}

	return nil
}

func (s *service) handleError(err error, logMessage string, logArgs ...interface{}) {
	s.deps.Log.Error(logMessage, err, logArgs)
	s.metrics.ErrorCount.Increment(err.Error())
}
