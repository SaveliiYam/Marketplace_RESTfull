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

func (r *CategoriesPostgres) Create(input marketplace.CategoriesList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", categoriesTable)

	row := tx.QueryRow(createCategoryQuery, input.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *CategoriesPostgres) GetById(id int) (marketplace.CategoriesList, error) {
	var category marketplace.CategoriesList

	query := fmt.Sprintf("SELECT id, title FROM %s WHERE id=$1", categoriesTable)
	err := r.db.Get(&category, query, id)
	return category, err
}

func (r *CategoriesPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", categoriesTable)
	_, err := r.db.Exec(query, id)
	return err
}

func (r *CategoriesPostgres) Update(id int, input marketplace.CategoriesList) error {
	query := fmt.Sprintf("UPDATE %s SET title=$1 WHERE id=$2", categoriesTable)
	_, err := r.db.Exec(query, input.Title, id)
	return err
}
