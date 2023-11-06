package service

import (
	"marketplace"
	"marketplace/pkg/repository"
)

type Authorization interface {
	CreateUser(marketplace.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	Update(id, brandId, categoriesId int, input marketplace.ProductList) error
	GetById(id int) (marketplace.ProductList, error)
}

type Service struct {
	Authorization
	Categories
	Brands
	Products
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoriesService(repos.Categories),
		Brands:        NewBrandService(repos.Brands),
		Products:      NewProductsService(repos.Products),
	}
}
