package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	ttlchecker "github.com/marisasha/ttl-check-app"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *ttlchecker.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username,password_hash) VALUES ($1, $2) RETURNING id", userTable)
	row := r.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, nil
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(username, password string) (ttlchecker.User, error) {

	var user ttlchecker.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err

}
