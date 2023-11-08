package repository

import (
	"fmt"
	"marketplace"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user marketplace.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash, status) values($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password, user.Status)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (marketplace.User, error) {
	var user marketplace.User
	query := fmt.Sprintf("SELECT id, status FROM %s WHERE username=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}

func (r *AuthPostgres) CheckStatus(userId int) (bool, error) {
	var status bool
	query := fmt.Sprintf("SELECT status FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&status, query, userId)
	return status, err
}
