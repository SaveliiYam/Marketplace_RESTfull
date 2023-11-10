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

func (s *CategoriesListServise) GetByString(input string) (int, error) {
	return s.repo.GetByString(input)
}

func (s *CategoriesListServise) CreateImage(id int, image_thumb *imageupload.Image) error {
	_, err := s.repo.GetById(id)
	if err != nil {
		return err
	}

	staticFolder := fmt.Sprintf("./static/categories/%d", id)
	if _, err := os.Stat(staticFolder); os.IsNotExist(err) {
		os.Mkdir(staticFolder, os.ModePerm)
	}

	name := fmt.Sprintf("./static/categories/%d/%d", id, time.Now().Unix())
	err = image_thumb.Save(name)
	return err
}

func (s *CategoriesListServise) GetImage(id int) (string, error) {
	_, err := s.repo.GetById(id)
	if err != nil {
		return "", err
	}

	imagePath := filepath.Join("./static/categories/", string(rune(id)))

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
