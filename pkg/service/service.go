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
}

type Brands interface {
	GetAllBrands() ([]marketplace.BrandsList, error)
	GetBrandById(id int) (marketplace.BrandsList, error)
}

type Service struct {
	Authorization
	Categories
	Brands
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoriesService(repos.Categories),
		Brands:        NewBrandService(repos.Brands),
	}
}
