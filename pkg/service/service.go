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

type Service struct {
	Authorization
	Categories
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Categories:    NewCategoriesService(repos.Categories),
	}
}
