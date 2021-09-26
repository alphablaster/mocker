package service

import (
	"github.com/alphablaster/mocker/pkg/entity"
	"github.com/alphablaster/mocker/pkg/repository"
)

type Authorization interface {
	
}

type Mock interface {
	Create(mock entity.Mock) (int, error)
}

type Header interface {
	Create(header entity.Header) (int, error)
}

type Service struct {
	Authorization
	Mock
	Header
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Mock: NewMockService(repos.Mock),
	}
}