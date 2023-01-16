package service

import (
	"context"

	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"app/internal/constants"
	"app/internal/logger"
	"app/internal/serviceB/domain"
	"app/internal/serviceB/repository"
	"app/internal/serviceB/service/metrics"
)

const (
	requestIDKey = "requestID"
	itemIDKey    = "itemID"
	itemObjKey   = "item"
	errorKey     = "error"
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
		s.handleError(ctx, err, FailedToGetAll, nil)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemB, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(ctx, err, FailedToParseUUID, logrus.Fields{requestIDKey: id})
		return nil, constants.ErrCreatingUUIDFromString
	}

	resp, err := s.deps.Repository.GetByID(ctx, itemID)
	if err != nil {
		s.handleError(ctx, err, FailedToGetByID, logrus.Fields{itemIDKey: itemID})
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemB) (*domain.ItemB, error) {
	resp, err := s.deps.Repository.Insert(ctx, item)
	if err != nil {
		s.handleError(ctx, err, FailedToCreate, logrus.Fields{itemObjKey: item})
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemB) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.handleError(ctx, err, FailedToParseUUID, logrus.Fields{requestIDKey: id})
		return constants.ErrCreatingUUIDFromString
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
		return constants.ErrCreatingUUIDFromString
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
