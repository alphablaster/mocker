package service

import (
	"github.com/alphablaster/mocker/pkg/entity"
	"github.com/alphablaster/mocker/pkg/repository"
)

type MockService struct {
	repo repository.Mock
}

func NewMockService(repo repository.Mock) *MockService {
	return &MockService{repo: repo}
}

func (s *MockService) Create(list entity.Mock) (int, error) {
	return s.repo.Create(list)
}

func (s *MockService) GetAll() ([]entity.Mock, error) {
	return s.repo.GetAll()
}