package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository/constants"
)

// TodoItemPostgres contains pointer on db instance
type TodoItemPostgres struct {
	db *sqlx.DB
}

// NewTodoItemPostgres returns a pointer on a new instance of TodoItemPostgres
func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

// Create creates a new model.TodoItem in db by specified list id. Returns id of new item model and error
func (r *TodoItemPostgres) Create(listId int, item model.TodoItem) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var itemId int
	createItemQuery := fmt.Sprintf("INSERT INTO %s (%s, %s, %s) VALUES ($1, $2, $3) RETURNING %s",
		constants.TodoItemsTable, constants.Title, constants.Description, constants.Done, constants.Id)
	row := tx.QueryRow(createItemQuery, item.Title, item.Description, item.Done)
	if err := row.Scan(&itemId); err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	createListsItemsQuery := fmt.Sprintf("INSERT INTO %s (%s, %s) VALUES ($1, $2) RETURNING %s",
		constants.ListsItemsTable, constants.ListId, constants.ItemId, constants.Id)
	_, err = tx.Exec(createListsItemsQuery, listId, itemId)
	if err != nil {
		_ = tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}

// GetAll returns error and slice of model.TodoItem by specified list id
func (r *TodoItemPostgres) GetAll(listId int) ([]model.TodoItem, error) {
	var items []model.TodoItem

	query := fmt.Sprintf(
		"SELECT ti.%s, ti.%s, ti.%s, ti.%s "+
			"FROM %s ti INNER JOIN %s li on ti.%s = li.%s "+
			"WHERE li.%s = $1",
		constants.Id, constants.Title, constants.Description, constants.Done,
		constants.TodoItemsTable, constants.ListsItemsTable, constants.Id, constants.ItemId,
		constants.ListId)
	err := r.db.Select(&items, query, listId)

	return items, err
}

// GetById returns error and model of model.TodoItem by specified list and item ids
func (r *TodoItemPostgres) GetById(listId, itemId int) (model.TodoItem, error) {
	var item model.TodoItem

	query := fmt.Sprintf(
		"SELECT ti.%s, ti.%s, ti.%s, ti.%s "+
			"FROM %s ti INNER JOIN %s li on ti.%s = li.%s "+
			"WHERE li.%s = $1 AND li.%s = $2",
		constants.Id, constants.Title, constants.Description, constants.Done,
		constants.TodoItemsTable, constants.ListsItemsTable, constants.Id, constants.ItemId,
		constants.ListId, constants.ItemId)
	err := r.db.Get(&item, query, listId, itemId)

	return item, err
}

// DeleteById deletes model.TodoItem from db by specified list and item ids. Returns an error if there is one
func (r *TodoItemPostgres) DeleteById(listId, itemId int) error {
	query := fmt.Sprintf(
		"DELETE FROM %s ti USING %s li "+
			"WHERE ti.%s = li.%s AND li.%s = $1 AND li.%s = $2",
		constants.TodoItemsTable, constants.ListsItemsTable,
		constants.Id, constants.ItemId, constants.ListId, constants.ItemId)

	_, err := r.db.Exec(query, listId, itemId)
	return err
}

// Update updates model.TodoItem in db by specified list, item ids and model of updated item. Returns an error if there is one
func (r *TodoItemPostgres) Update(listId int, itemId int, item model.UpdateItemInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if item.Title != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", constants.Title, argId))
		args = append(args, *item.Title)
		argId++
	}

	if item.Description != nil {
		setValues = append(setValues, fmt.Sprintf("%s=$%d", constants.Description, argId))
		args = append(args, *item.Description)
		argId++
	}

	setValues = append(setValues, fmt.Sprintf("%s=$%d", constants.Done, argId))
	args = append(args, item.Done)
	argId++

	args = append(args, listId, itemId)
	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf("UPDATE %s ti SET %s FROM %s li "+
		"WHERE ti.%s=li.%s AND li.%s=$%d AND li.%s=$%d",
		constants.TodoItemsTable, setQuery, constants.ListsItemsTable,
		constants.Id, constants.ItemId, constants.ListId, argId, constants.ItemId, argId+1)

	_, err := r.db.Exec(query, args...)
	return err
}
