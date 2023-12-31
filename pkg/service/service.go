package service

import (
	"marketplace"
	"marketplace/pkg/repository"

	"github.com/olahol/go-imageupload"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(marketplace.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	CheckStatus(userId int) (bool, error)
}
type Categories interface {
	GetAllCategories() ([]marketplace.CategoriesList, error)
	Create(input marketplace.CategoriesList) (int, error)
	GetById(id int) (marketplace.CategoriesList, error)
	Delete(id int) error
	Update(id int, input marketplace.CategoriesList) error
	GetByString(input string) (int, error)
	CreateImage(id int, image_thumb *imageupload.Image) error
	GetImage(id int) (string, error)
}

type Brands interface {
	GetAllBrands() ([]marketplace.BrandsList, error)
	GetById(id int) (marketplace.BrandsList, error)
	Create(input marketplace.BrandsList) (int, error)
	Delete(id int) error
	Update(id int, input marketplace.BrandsList) error
	GetByString(input string) (int, error)
	CreateImage(id int, image_thumb *imageupload.Image) error
	GetImage(id int) (string, error)
}

type Products interface {
	Create(input marketplace.ProductList, brandId, categoryId int) (int, error)
	GetAll() ([]marketplace.ProductList, error)
	Delete(id int) error
	Update(id, brandId, categoriesId int, input marketplace.ProductList) error
	GetById(id int) (marketplace.ProductList, error)
	CreateImage(id int, image_thumb *imageupload.Image) error
	GetImage(id int) (string, error)
}

type Basket interface {
	GetAll(id int) ([]marketplace.BusketList, error)
	Create(userId int, input marketplace.BusketList) (int, error)
	GetById(userId, basketId int) (marketplace.BusketList, error)
	Delete(userId, basketId int) error
}

type Service struct {
	Authorization
	Categories
	Brands
	Products
	Basket
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoriesService(repos.Categories),
		Brands:        NewBrandService(repos.Brands),
		Products:      NewProductsService(repos.Products),
		Basket:        NewBasketService(repos.Basket),
	}
}
