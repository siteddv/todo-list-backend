package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"todolistBackend/internal/model"
)

// AuthPostgres contains pointer on db instance
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres returns a pointer on a new instance of AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

// CreateUser creates model.User in DB using specified user model. It returns new user id and error
func (r *AuthPostgres) CreateUser(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) values ($1, $2, $3) RETURNING %s",
		UsersTable, Name, Username, PasswordHash, Id)

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
		Id, UsersTable, Username, PasswordHash)

	err := r.db.Get(&user, query, username, password)

	return user, err
}
