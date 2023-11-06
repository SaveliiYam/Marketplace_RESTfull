package service

import (
	"marketplace"
	"marketplace/pkg/repository"
)

type BrandService struct {
	repo repository.Brands
}

func NewBrandService(repo repository.Brands) *BrandService {
	return &BrandService{repo: repo}
}

func (s *BrandService) GetAllBrands() ([]marketplace.BrandsList, error) {
	return s.repo.GetAllBrands()
}

func (s *BrandService) GetById(id int) (marketplace.BrandsList, error) {
	return s.repo.GetById(id)
}

func (s *BrandService) Create(input marketplace.BrandsList) (int, error) {
	return s.repo.Create(input)
}
func (s *BrandService) Delete(id int) error {
	return s.repo.Delete(id)
}
func (s *BrandService) Update(id int, input marketplace.BrandsList) error {
	return s.repo.Update(id, input)
}

func (s *BrandService) GetByString(input string) (int, error) {
	return s.repo.GetByString(input)
}
