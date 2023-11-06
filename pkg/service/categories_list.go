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

func (s *CategoriesListServise) Create(input marketplace.CategoriesList) (int, error) {
	return s.repo.Create(input)
}

func (s *CategoriesListServise) GetById(id int) (marketplace.CategoriesList, error) {
	return s.repo.GetById(id)
}

func (s *CategoriesListServise) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *CategoriesListServise) Update(id int, input marketplace.CategoriesList) error {
	return s.repo.Update(id, input)
}
