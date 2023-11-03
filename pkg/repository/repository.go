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
type Shoes interface {
	GetAllShoes() ([]marketplace.ProductList, error)
}

type Brands interface {
	GetAllBrands() ([]marketplace.BrandsList, error)
}

type Repository struct {
	Authorization
	Categories
	Shoes
	Brands
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Categories:    NewCategoryPostgres(db),
		Shoes:         NewShoesPostgres(db),
		Brands:        NewBrandsPostgres(db),
	}
}
