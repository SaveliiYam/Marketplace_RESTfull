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
	query := fmt.Sprintf("SELECT id, title, description FROM %s", brandsTable)
	err := r.db.Select(&brands, query)
	return brands, err
}

func (r *BrandsListRepository) GetById(id int) (marketplace.BrandsList, error) {
	var brand marketplace.BrandsList

	query := fmt.Sprintf("SELECT id, title, description FROM %s WHERE id=$1", brandsTable)
	err := r.db.Get(&brand, query, id)
	return brand, err
}

func (r *BrandsListRepository) Create(input marketplace.BrandsList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBrandQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", brandsTable)

	row := tx.QueryRow(createBrandQuery, input.Title, input.Description)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *BrandsListRepository) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", brandsTable)
	_, err := r.db.Exec(query, id)
	return err
}
func (r *BrandsListRepository) Update(id int, input marketplace.BrandsList) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2 WHERE id=$3", brandsTable)
	_, err := r.db.Exec(query, input.Title, input.Description, id)
	return err
}

func (r *BrandsListRepository) GetByString(input string) (int, error) {
	var brand marketplace.BrandsList

	query := fmt.Sprintf("SELECT id, title FROM %s WHERE title=$1", brandsTable)
	err := r.db.Get(&brand, query, input)
	return brand.Id, err
}
