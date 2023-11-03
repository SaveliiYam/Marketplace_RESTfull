package service

import (
	"marketplace"
	"marketplace/pkg/repository"
)

type CategoriesListServise struct {
	repo repository.Categories
}

func NewCategoriesService(repo repository.Categories) *CategoriesListServise {
	return &CategoriesListServise{repo: repo}
}

func (s *CategoriesListServise) GetAllCategories() ([]marketplace.CategoriesList, error) {
	return s.repo.GetAllCategories()
}
