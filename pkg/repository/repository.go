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
	Create(input marketplace.CategoriesList) (int, error)
	GetById(id int) (marketplace.CategoriesList, error)
	Delete(id int) error
	Update(id int, input marketplace.CategoriesList) error
}

type Brands interface {
	GetAllBrands() ([]marketplace.BrandsList, error)
	GetById(id int) (marketplace.BrandsList, error)
	Create(input marketplace.BrandsList) (int, error)
	Delete(id int) error
	Update(id int, input marketplace.BrandsList) error
}

type Repository struct {
	Authorization
	Categories
	Brands
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Categories:    NewCategoryPostgres(db),
		Brands:        NewBrandsPostgres(db),
	}
}
