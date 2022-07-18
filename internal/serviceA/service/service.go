package service

import (
	uuid "github.com/satori/go.uuid"

	"microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/pkg"
)

type Service interface {
	GetAll() ([]domain.ItemA, error)
	GetOneByID(id uuid.UUID) (domain.ItemA, error)
	Create(item domain.ItemA) (domain.ItemA, error)
	Update(id uuid.UUID, item domain.ItemA) error
	Delete(id uuid.UUID) error
}

func New(log pkg.Logger) Service {
	return service{
		log: log,
	}
}

type service struct {
	log pkg.Logger
}

func (s service) GetAll() ([]domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) GetOneByID(id uuid.UUID) (domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(item domain.ItemA) (domain.ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(id uuid.UUID, item domain.ItemA) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
