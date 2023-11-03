package service

import (
	"marketplace"
	"marketplace/pkg/repository"
)

type ShoesService struct {
	repo repository.Shoes
}

func NewShoesService(repo repository.Shoes) *ShoesService {
	return &ShoesService{repo: repo}
}

func (s *ShoesService) GetAllShoes() ([]marketplace.ProductList, error) {
	return s.repo.GetAllShoes()
}
