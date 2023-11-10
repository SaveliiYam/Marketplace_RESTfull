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

func (s *BrandService) CreateImage(id int, image_thumb *imageupload.Image) error {
	_, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	staticFolder := fmt.Sprintf("./static/brands/%d", id)
	if _, err := os.Stat(staticFolder); os.IsNotExist(err) {
		os.Mkdir(staticFolder, os.ModePerm)
	}

	name := fmt.Sprintf("./static/brands/%d/%d", id, time.Now().Unix())
	err = image_thumb.Save(name)
	return err
}

func (s *BrandService) GetImage(id int) (string, error) {
	_, err := s.repo.GetById(id)
	if err != nil {
		return "", err
	}

	imagePath := filepath.Join("./static/brands/", string(rune(id)))

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
