package service

import (
	"context"

	"microservices-boilerplate/internal/pkg"
	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/internal/serviceA/repository"
)

type Service interface {
	GetAll(ctx context.Context) ([]*domain.ItemA, error)
	GetOneByID(ctx context.Context, id string) (*domain.ItemA, error)
	Create(ctx context.Context, item domain.ItemA) (*domain.ItemA, error)
	Update(ctx context.Context, id string, item domain.ItemA) error
	Delete(ctx context.Context, id string) error
}

func New(log pkg.Logger, repo repository.Repository) Service {
	return service{
		log:        log,
		repository: repo,
	}
}

type service struct {
	log        pkg.Logger
	repository repository.Repository
}

func (s service) GetAll(ctx context.Context) ([]*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetOneByID(ctx context.Context, id string) (*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(ctx context.Context, item domain.ItemA) (*domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(ctx context.Context, id string, item domain.ItemA) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
