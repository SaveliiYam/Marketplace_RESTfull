package repository

import (
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(marketplace.User) (int, error)
	GetUser(username, password string) (marketplace.User, error)
}
type Categories interface {
	GetAllCategories() ([]marketplace.CategoriesList, error)
}

type Repository struct {
	Authorization
	Categories
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Categories:    NewCategoryPostgres(db),
	}
}
