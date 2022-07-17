package service

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/mock"
	"microservices-boilerplate/internal/serviceB/domain"
)

type Service struct {
	mock.Mock
}

func (s *Service) Get() ([]domain.ItemB, error) {
	called := s.Called()
	return called.Get(0).([]domain.ItemB), called.Error(1)
}

func (s *Service) Find(id uuid.UUID) (domain.ItemB, error) {
	called := s.Called(id)
	return called.Get(0).(domain.ItemB), called.Error(1)
}

func (s *Service) Create(item domain.ItemB) (domain.ItemB, error) {
	called := s.Called(item)
	return called.Get(0).(domain.ItemB), called.Error(1)
}

func (s *Service) Update(id uuid.UUID, item domain.ItemB) error {
	called := s.Called(id, item)
	return called.Error(0)
}

func (s *Service) Delete(id uuid.UUID) error {
	called := s.Called(id)
	return called.Error(0)
}
