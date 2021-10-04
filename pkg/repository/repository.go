package repository

import (
	"github.com/alphablaster/mocker/pkg/entity"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {

}

type Mock interface {
	Create(mock entity.Mock) (int, error)
	GetAll() ([]entity.Mock, error)
	Delete(listId int) error
}

type Header interface {
	Create(header entity.Header) (int, error)
}

type Repository struct {
	Authorization
	Mock
	Header
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Mock: NewMockPostgres(db),
	}
}
