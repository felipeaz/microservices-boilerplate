package service

import (
	uuid "github.com/satori/go.uuid"
	. "microservices-boilerplate/internal/serviceB/domain"
	"microservices-boilerplate/pkg"
)

type Service interface {
	Get() ([]ItemB, error)
	Find(id uuid.UUID) (ItemB, error)
	Create(item ItemB) (ItemB, error)
	Update(id uuid.UUID, item ItemB) error
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

func (s service) Get() ([]ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Find(id uuid.UUID) (ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Create(item ItemB) (ItemB, error) {
	//TODO implement me
	panic("implement me")
}

func (s service) Update(id uuid.UUID, item ItemB) error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete(id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}
