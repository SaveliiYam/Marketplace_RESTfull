package service

import (
	"errors"
	"fmt"
	"marketplace"
	"marketplace/pkg/repository"
	"os"
	"path/filepath"
	"time"

	"github.com/olahol/go-imageupload"
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

func (s *ProductsListService) CreateImage(id int, image_thumb *imageupload.Image) error {
	_, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	staticFolder := fmt.Sprintf("./static/products/%d", id)
	if _, err := os.Stat(staticFolder); os.IsNotExist(err) {
		os.Mkdir(staticFolder, os.ModePerm)
	}

	name := fmt.Sprintf("./static/products/%d/%d", id, time.Now().Unix())
	err = image_thumb.Save(name)
	return err
}

func (s *ProductsListService) GetImage(id int) (string, error) {
	_, err := s.repo.GetById(id)
	if err != nil {
		return "", err
	}

	imagePath := filepath.Join("./static/products/", string(rune(id)))

	dir, err := os.Open(imagePath)
	if err != nil {
		return "", err
	}
	defer dir.Close()

	files, err := dir.Readdirnames(-1)
	if err != nil {
		return "", err
	}

	if len(files) > 0 {
		imageName := files[0]
		imagePath := filepath.Join(imagePath, imageName)
		return imagePath, nil
	}
	return "", errors.New("no images found in the specified folder")
}
