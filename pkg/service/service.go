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
}

type Shoes interface {
	GetAllShoes() ([]marketplace.ProductList, error)
}

type Brands interface {
	GetAllBrands() ([]marketplace.BrandsList, error)
}

type Service struct {
	Authorization
	Categories
	Shoes
	Brands
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoriesService(repos.Categories),
		Shoes:         NewShoesService(repos.Shoes),
		Brands:        NewBrandService(repos.Brands),
	}
}
