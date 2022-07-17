package service

import (
	uuid "github.com/satori/go.uuid"
	"microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/pkg"
)

type Service interface {
	Get() ([]domain.ItemB, error)
	Find(id uuid.UUID) (domain.ItemB, error)
	Create(item domain.ItemB) (domain.ItemB, error)
	Update(id uuid.UUID, item domain.ItemB) error
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

func (s service) Get() ([]domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Find(id uuid.UUID) (domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(item domain.ItemB) (domain.ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(id uuid.UUID, item domain.ItemB) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
