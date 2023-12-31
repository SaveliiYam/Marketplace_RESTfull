package repository

import (
	"fmt"
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type BasketListRepository struct {
	db *sqlx.DB
}

func NewBasketPostgres(db *sqlx.DB) *BasketListRepository {
	return &BasketListRepository{db: db}
}

func (r *BasketListRepository) GetAll(id int) ([]marketplace.BusketList, error) {
	var baskets []marketplace.BusketList
	query := fmt.Sprintf("SELECT id, user_id, product_id FROM %s WHERE user_id=$1", basketsTable)
	err := r.db.Select(&baskets, query, id)
	return baskets, err
}

func (r *BasketListRepository) Create(userId int, input marketplace.BusketList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createBasketQuery := fmt.Sprintf("INSERT INTO %s (user_id, product_id) VALUES ($1, $2) RETURNING id", basketsTable)

	row := tx.QueryRow(createBasketQuery, userId, input.ProductId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *BasketListRepository) GetById(userId, basketId int) (marketplace.BusketList, error) {
	var basket marketplace.BusketList

	query := fmt.Sprintf("SELECT id, user_id, product_id FROM %s WHERE user_id=$1 AND id=$2", basketsTable)
	err := r.db.Get(&basket, query, userId, basketId)
	return basket, err
}
func (r *BasketListRepository) Delete(userId, basketId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1 AND user_id=$2", basketsTable)
	_, err := r.db.Exec(query, basketId, userId)
	return err
}
