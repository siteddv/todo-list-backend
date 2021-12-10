package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository/constants"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser creates model.User in DB using specified user model. It returns new user id and error
func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) values ($1, $2, $3) RETURNING id",
		constants.UsersTable, constants.Name, constants.Username, constants.PasswordHash)

	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

// GetUser returns error and model.User by specified user username, password
func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT %s FROM %s WHERE %s=$1 AND %s=$2",
		constants.Id, constants.UsersTable, constants.Username, constants.PasswordHash)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
