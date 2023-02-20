package service

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"app/internal/errors"
	"app/internal/logger"
	"app/internal/serviceA/domain"
	"app/internal/serviceA/repository"
	"app/internal/serviceA/service/metrics"
)

const (
	requestIDKey = "requestID"
	itemIDKey    = "itemID"
	itemObjKey   = "item"
	errorKey     = "error"
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
		s.handleError(ctx, err, FailedToGetAll, nil)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemA, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(ctx, err, FailedToParseUUID, logrus.Fields{requestIDKey: id})
		return nil, errors.ErrCreatingUUIDFromString
	}

	resp, err := s.deps.Repository.GetByID(ctx, itemID)
	if err != nil {
		s.handleError(ctx, err, FailedToGetByID, logrus.Fields{itemIDKey: itemID})
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	resp, err := s.deps.Repository.Insert(ctx, item)
	if err != nil {
		s.handleError(ctx, err, FailedToCreate, logrus.Fields{itemObjKey: item})
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemA) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(ctx, err, FailedToParseUUID, logrus.Fields{requestIDKey: id})
		return errors.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Update(ctx, itemID, item); err != nil {
		s.handleError(ctx, err, FailedToUpdate, logrus.Fields{itemIDKey: itemID, itemObjKey: item})
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(ctx, err, FailedToParseUUID, logrus.Fields{requestIDKey: id})
		return errors.ErrCreatingUUIDFromString
	}

	if err = s.deps.Repository.Remove(ctx, itemID); err != nil {
		s.handleError(ctx, err, FailedToDelete, logrus.Fields{itemIDKey: itemID})
		return err
	}

	return nil
}

func (s *service) handleError(ctx context.Context, err error, logMessage string, fields logrus.Fields) {
	s.deps.Log.Error(ctx, err, logMessage, fields)
	s.metrics.ErrorCount.Increment(errorKey, err.Error())
}
