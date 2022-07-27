package service

import (
	"context"

	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/errors"
	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/serviceA/repository"
)

type Service interface {
	GetAll(ctx context.Context) ([]*domain.ItemA, error)
	GetOneByID(ctx context.Context, id string) (*domain.ItemA, error)
	Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id string, item *domain.ItemA) error
	Delete(ctx context.Context, id string) error
}

func New(log pkg.Logger, repo repository.Repository) Service {
	return &service{
		log:        log,
		repository: repo,
	}
}

type service struct {
	log        pkg.Logger
	repository repository.Repository
}

func (s *service) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	resp, err := s.repository.GetAll(ctx)
	if err != nil {
		s.log.Error("failed to get all item a", err)
		return nil, err
	}

	return resp, nil
}

func (s *service) GetOneByID(ctx context.Context, id string) (*domain.ItemA, error) {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.log.Error("failed to parse id to UUID", err)
		return nil, errors.ErrCreatingUUIDFromString
	}

	resp, err := s.repository.GetByID(ctx, itemID)
	if err != nil {
		s.log.Error("failed to get item with id", itemID, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Create(ctx context.Context, item *domain.ItemA) (*domain.ItemA, error) {
	resp, err := s.repository.Insert(ctx, item)
	if err != nil {
		s.log.Error("failed to create item", item, err)
		return nil, err
	}

	return resp, nil
}

func (s *service) Update(ctx context.Context, id string, item *domain.ItemA) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.log.Error("failed to parse id to UUID", err)
		return errors.ErrCreatingUUIDFromString
	}

	if err = s.repository.Update(ctx, itemID, item); err != nil {
		s.log.Error("failed to update item", itemID, item, err)
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, id string) error {
	itemID, err := uuid.FromString(id)
	if err != nil {
		s.log.Error("failed to parse id to UUID", err)
		return errors.ErrCreatingUUIDFromString
	}

	if err = s.repository.Remove(ctx, itemID); err != nil {
		s.log.Error("failed to delete item", itemID, err)
		return err
	}

	return nil
}
