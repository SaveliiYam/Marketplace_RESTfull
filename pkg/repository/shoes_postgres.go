package repository

import (
	"fmt"
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type ShoesPostgres struct {
	db *sqlx.DB
}

func NewShoesPostgres(db *sqlx.DB) *ShoesPostgres {
	return &ShoesPostgres{db: db}
}

func (r *ShoesPostgres) GetAllShoes() ([]marketplace.ProductList, error) {
	var shoes []marketplace.ProductList
	query := fmt.Sprintf("SELECT id, title FROM %s", shoesTable)
	err := r.db.Select(&shoes, query)
	return shoes, err
}
