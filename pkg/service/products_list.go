package service

import (
	"marketplace"
	"marketplace/pkg/repository"
)

type ProductsListService struct {
	repo repository.Products
}

func NewProductsService(repo repository.Products) *ProductsListService {
	return &ProductsListService{repo: repo}
}

func (p *ProductsListService) Create(input marketplace.ProductList, brandId, categoryId int) (int, error) {
	return p.repo.Create(input, brandId, categoryId)
}
func (p *ProductsListService) GetAll() ([]marketplace.ProductList, error) {
	return p.repo.GetAll()
}

func (p *ProductsListService) GetById(id int) (marketplace.ProductList, error) {
	return p.repo.GetById(id)
}
func (p *ProductsListService) Delete(id int) error {
	return p.repo.Delete(id)
}

func (p *ProductsListService) Update(id, brandId, categoriesId int, input marketplace.ProductList) error {
	return p.repo.Update(id, brandId, categoriesId, input)
}
