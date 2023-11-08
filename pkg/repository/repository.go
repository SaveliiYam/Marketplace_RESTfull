package repository

import (
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(marketplace.User) (int, error)
	GetUser(username, password string) (marketplace.User, error)
	CheckStatus(userId int) (bool, error)
}
type Categories interface {
	GetAllCategories() ([]marketplace.CategoriesList, error)
	Create(input marketplace.CategoriesList) (int, error)
	GetById(id int) (marketplace.CategoriesList, error)
	Delete(id int) error
	Update(id int, input marketplace.CategoriesList) error
	GetByString(input string) (int, error)
}

type Brands interface {
	GetAllBrands() ([]marketplace.BrandsList, error)
	GetById(id int) (marketplace.BrandsList, error)
	Create(input marketplace.BrandsList) (int, error)
	Delete(id int) error
	Update(id int, input marketplace.BrandsList) error
	GetByString(input string) (int, error)
}

type Products interface {
	Create(input marketplace.ProductList, brandId, categoryId int) (int, error)
	GetAll() ([]marketplace.ProductList, error)
	Delete(id int) error
	Update(id, brandId, categoryId int, input marketplace.ProductList) error
	GetById(id int) (marketplace.ProductList, error)
}

type Basket interface {
	GetAll(id int) ([]marketplace.BusketList, error)
	Create(userId int, input marketplace.BusketList) (int, error)
	GetById(userId, basketId int) (marketplace.BusketList, error)
	Delete(userId, basketId int) error
}

type Repository struct {
	Authorization
	Categories
	Brands
	Products
	Basket
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Categories:    NewCategoryPostgres(db),
		Brands:        NewBrandsPostgres(db),
		Products:      NewProductsPostgres(db),
		Basket:        NewBasketPostgres(db),
	}
}
