package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"todolistBackend/internal/model"
)

// TodoListPostgres contains pointer on db instance
type TodoListPostgres struct {
	db *sqlx.DB
}

// NewTodoListPostgres returns a pointer on a new instance of TodoListPostgres
func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

// Create creates a new model.TodoList in db by specified user id. Returns id of new list and error
func (r *TodoListPostgres) Create(userId int, list model.TodoList) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var listId int
	createListQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING %s",
		TodoListTable, Title, Description, Id)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)
	if err := row.Scan(&listId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createUsersListsQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING %s",
		UsersListsTable, UserId, ListId, Id)
	_, err = tx.Exec(createUsersListsQuery, userId, listId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

// GetAll returns error and slice of model.TodoList by specified user id
func (r *TodoListPostgres) GetAll(userId int) ([]model.TodoList, error) {
	var lists []model.TodoList

	query := fmt.Sprintf(
		"SELECT tl.%s, tl.%s, tl.%s "+
			"FROM %s tl INNER JOIN %s ul on tl.%s = ul.%s "+
			"WHERE ul.%s = $1",
		Id, Title, Description,
		TodoListTable, UsersListsTable, Id, ListId,
		UserId)
	err := r.db.Select(&lists, query, userId)

	return lists, err
}

// GetById returns error and model of model.TodoList by specified user and list ids
func (r *TodoListPostgres) GetById(userId, listId int) (model.TodoList, error) {
	var list model.TodoList

	query := fmt.Sprintf(
		"SELECT tl.%s, tl.%s, tl.%s "+
			"FROM %s tl INNER JOIN %s ul on tl.%s = ul.%s "+
			"WHERE ul.%s = $1 AND ul.%s = $2",
		Id, Title, Description,
		TodoListTable, UsersListsTable, Id, ListId,
		UserId, ListId)
	err := r.db.Get(&list, query, userId, listId)

	return list, err
}

// DeleteById deletes model.TodoList from db by specified list and user ids. Returns an error if there is one.
// This method also deletes all the items from specified list
func (r *TodoListPostgres) DeleteById(userId, listId int) error {
	var intListItems []int

	selectListItemsQuery := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s=$1",
		Id, ListsItemsTable, ListId)
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
		TodoItemsTable, Id, joinedListItems)
	_, err = r.db.Exec(deleteListItemsQuery)
	if err != nil {
		return err
	}

	deleteListQuery := fmt.Sprintf(
		"DELETE FROM %s tl USING %s ul "+
			"WHERE tl.%s = ul.%s AND ul.%s = $1 AND ul.%s = $2",
		TodoListTable, UsersListsTable,
		Id, ListId, UserId, ListId)

	_, err = r.db.Exec(deleteListQuery, userId, listId)
	return err
}

// Update updates model.TodoList in db by specified list, user ids and model of updated item. Returns an error if there is one
func (r *TodoListPostgres) Update(userId, listId int, list model.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if list.Title != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", Title, argId))
		args = append(args, *list.Title)
		argId++
	}

	if list.Description != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", Description, argId))
		args = append(args, *list.Description)
		argId++
	}

	args = append(args, listId, userId)
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul "+
		"WHERE tl.%s=ul.%s AND ul.%s=$%d AND ul.%s=$%d",
		TodoListTable, setQuery, UsersListsTable,
		Id, ListId, ListId, argId, UserId, argId+1)

	_, err := r.db.Exec(query, args...)
	return err
}
