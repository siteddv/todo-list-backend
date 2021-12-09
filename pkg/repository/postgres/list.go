package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
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
		_ = tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING %s",
		constants.UsersListsTable, constants.UserId, constants.ListId, constants.Id)
	_, err = tx.Exec(createUsersListsQuery, userId, listId)
	if err != nil {
		_ = tx.Rollback()
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

func (r *TodoListPostgres) DeleteById(userId, listId int) error {
	var intListItems []int

	selectListItemsQuery := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s=$1",
		constants.Id, constants.ListsItemsTable, constants.ListId)
	err := r.db.Select(&intListItems, selectListItemsQuery, listId)
	if err != nil {
		return err
	}

	stringListItems := make([]string, len(intListItems))
	for i, s := range intListItems {
		stringListItems[i] = strconv.Itoa(s)
	}
	joinedListItems := strings.Join(stringListItems, ", ")

	deleteListItemsQuery := fmt.Sprintf(
		"DELETE FROM %s WHERE %s IN (%s)",
		constants.TodoItemsTable, constants.Id, joinedListItems)
	_, err = r.db.Exec(deleteListItemsQuery)
	if err != nil {
		return err
	}

	deleteListQuery := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul "+
			"WHERE tl.%s = ul.%s AND ul.%s = $1 AND ul.%s = $2",
		constants.TodoListTable, constants.UsersListsTable,
		constants.Id, constants.ListId, constants.UserId, constants.ListId)

	_, err = r.db.Exec(deleteListQuery, userId, listId)
	return err
}

func (r *TodoListPostgres) Update(userId, listId int, list todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if list.Title != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", constants.Title, argId))
		args = append(args, *list.Title)
		argId++
	}

	if list.Description != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", constants.Description, argId))
		args = append(args, *list.Description)
		argId++
	}

	args = append(args, listId, userId)
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul "+
		"WHERE tl.%s=ul.%s AND ul.%s=$%d AND ul.%s=$%d",
		constants.TodoListTable, setQuery, constants.UsersListsTable,
		constants.Id, constants.ListId, constants.ListId, argId, constants.UserId, argId+1)

	_, err := r.db.Exec(query, args...)
	return err
}
