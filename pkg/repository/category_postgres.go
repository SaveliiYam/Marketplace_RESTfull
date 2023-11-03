package repository

import (
	"fmt"
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type CategoriesPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoriesPostgres {
	return &CategoriesPostgres{db: db}
}

func (r *CategoriesPostgres) GetAllCategories() ([]marketplace.CategoriesList, error) {
	var categories []marketplace.CategoriesList
	query := fmt.Sprintf("SELECT id, title FROM %s", categoriesTable)
	err := r.db.Select(&categories, query)
	return categories, err
}
