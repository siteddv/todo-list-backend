package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "todolistBackend"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) values ($1, $2, $3) RETURNING id",
		usersTable, name, username, password_hash)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(userName, password string) (todo.User, error) {
	var user todo.User
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1 AND %s=$2",
		id, usersTable, username, password_hash)

	err := r.db.Get(&user, query, userName, password)

	return user, err
}
