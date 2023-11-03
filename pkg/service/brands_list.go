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
