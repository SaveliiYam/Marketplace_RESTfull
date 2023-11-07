package service

import (
	"marketplace"
	"marketplace/pkg/repository"
)

type BasketService struct {
	repo repository.Basket
}

func NewBasketService(repo repository.Basket) *BasketService {
	return &BasketService{repo: repo}
}

func (s *BasketService) GetAll(id int) ([]marketplace.BusketList, error) {
	return s.repo.GetAll(id)
}
func (s *BasketService) Create(userId int, input marketplace.BusketList) (int, error) {
	return s.repo.Create(userId, input)
}
