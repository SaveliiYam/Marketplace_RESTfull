package repository

import (
	"fmt"
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type ProductsPostgres struct {
	db *sqlx.DB
}

func NewProductsPostgres(db *sqlx.DB) *ProductsPostgres {
	return &ProductsPostgres{db: db}
}

func (r *ProductsPostgres) Create(input marketplace.ProductList, brandId, categoryId int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createProductQuery := fmt.Sprintf("Insert INTO %s (title, description, price, brand_id, categories_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", productsTable)

	row := tx.QueryRow(createProductQuery, input.Title, input.Description, input.Price, brandId, categoryId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}
func (r *ProductsPostgres) GetAll() ([]marketplace.ProductList, error) {
	var products []marketplace.ProductList
	query := fmt.Sprintf("SELECT id, title, description, price, brand_id, categories_id FROM %s", productsTable)
	err := r.db.Select(&products, query)
	return products, err
}

func (r *ProductsPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", productsTable)
	_, err := r.db.Exec(query, id)
	return err
}
func (r *ProductsPostgres) Update(id, brandId, categoryId int, input marketplace.ProductList) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1, description=$2, price=$3, brand_id=$4, categories_id=$5 WHERE id=$6", productsTable)
	_, err := r.db.Exec(query, input.Title, input.Description, input.Price, brandId, categoryId, id)
	return err
}
func (r *ProductsPostgres) GetById(id int) (marketplace.ProductList, error) {
	var product marketplace.ProductList

	query := fmt.Sprintf("SELECT id, title, description, price, brand_id, categories_id FROM %s WHERE id=$1", productsTable)
	err := r.db.Get(&product, query, id)
	return product, err
}
