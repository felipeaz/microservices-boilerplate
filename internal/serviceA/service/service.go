package service

import (
	uuid "github.com/satori/go.uuid"
	. "microservices-boilerplate/internal/serviceA/domain"
	"microservices-boilerplate/pkg"
)

type Service interface {
	Get() ([]ItemA, error)
	Find(id uuid.UUID) (ItemA, error)
	Create(item ItemA) (ItemA, error)
	Update(id uuid.UUID, item ItemA) error
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

func (s service) Get() ([]ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Find(id uuid.UUID) (ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(item ItemA) (ItemA, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(id uuid.UUID, item ItemA) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
