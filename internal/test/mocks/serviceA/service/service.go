package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"

	"microservices-boilerplate/internal/serviceA/domain"
)

type Service struct {
	mock.Mock
}

func (s *Service) GetAll() ([]domain.ItemA, error) {
	called := s.Called()
	return called.Get(0).([]domain.ItemA), called.Error(1)
}

func (s *Service) GetOneByID(id uuid.UUID) (domain.ItemA, error) {
	called := s.Called(id)
	return called.Get(0).(domain.ItemA), called.Error(1)
}

func (s *Service) Create(item domain.ItemA) (domain.ItemA, error) {
	called := s.Called(item)
	return called.Get(0).(domain.ItemA), called.Error(1)
}

func (s *Service) Update(id uuid.UUID, item domain.ItemA) error {
	called := s.Called(id, item)
	return called.Error(0)
}

func (s *Service) Delete(id uuid.UUID) error {
	called := s.Called(id)
	return called.Error(0)
}
