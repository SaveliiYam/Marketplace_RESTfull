package repository

import (
	"fmt"
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type BrandsListRepository struct {
	db *sqlx.DB
}

func NewBrandsPostgres(db *sqlx.DB) *BrandsListRepository {
	return &BrandsListRepository{db: db}
}

func (r *BrandsListRepository) GetAllBrands() ([]marketplace.BrandsList, error) {
	var brands []marketplace.BrandsList
	query := fmt.Sprintf("SELECT id, title FROM %s", brandsTable)
	err := r.db.Select(&brands, query)
	return brands, err
}

func (r *BrandsListRepository) GetById(id int) (marketplace.BrandsList, error) {
	var brand marketplace.BrandsList

	query := fmt.Sprintf("SELECT id, title, description FROM %s", brandsTable)
	err := r.db.Get(&brand, query)
	return brand, err
}
