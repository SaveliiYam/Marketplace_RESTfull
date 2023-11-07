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

func (s *BasketService) GetById(userId, basketId int) (marketplace.BusketList, error) {
	return s.repo.GetById(userId, basketId)
}
func (s *BasketService) Delete(userId, basketId int) error {
	return s.repo.Delete(userId, basketId)
}
