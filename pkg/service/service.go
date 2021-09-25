package service

import "github.com/alphablaster/mocker/pkg/repository"

type Authorization interface {
	
}

type Mock interface {
	
}

type Service struct {
	Authorization
	Mock
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}