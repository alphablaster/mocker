package repository

type Authorization interface {

}

type Mock interface {

}

type Repository struct {
	Authorization
	Mock
}

func NewRepository() *Repository {
	return &Repository{}
}
