package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {

}

type Mock interface {

}

type Repository struct {
	Authorization
	Mock
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
