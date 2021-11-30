package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	todo "todolistBackend"
	"todolistBackend/pkg/repository/constants"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, list todo.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING %s",
		constants.TodoListTable, constants.Title, constants.Description, constants.Id)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING %s",
		constants.UsersListsTable, constants.UserId, constants.ListId, constants.Id)
	_, err = tx.Exec(createUsersListsQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}
