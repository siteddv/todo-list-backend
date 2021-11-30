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

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf(
		"SELECT tl.%s, tl.%s, tl.%s "+
			"FROM %s tl INNER JOIN %s ul on tl.%s = ul.%s "+
			"WHERE ul.%s = $1",
		constants.Id, constants.Title, constants.Description,
		constants.TodoListTable, constants.UsersListsTable, constants.Id, constants.ListId,
		constants.UserId)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

func (r *TodoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(
		"SELECT tl.%s, tl.%s, tl.%s "+
			"FROM %s tl INNER JOIN %s ul on tl.%s = ul.%s "+
			"WHERE ul.%s = $1 AND ul.%s = $2",
		constants.Id, constants.Title, constants.Description,
		constants.TodoListTable, constants.UsersListsTable, constants.Id, constants.ListId,
		constants.UserId, constants.ListId)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}
