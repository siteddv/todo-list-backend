package repository

import (
	"github.com/jmoiron/sqlx"
	"todolistBackend/pkg/model"
	postgres "todolistBackend/pkg/repository/postgres"
)

type Authorization interface {
	// CreateUser creates model.User in DB using specified user model. It returns new user id and error
	CreateUser(user model.User) (int, error)

	// GetUser returns error and model.User by specified user username, password
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	// Create creates a new model.TodoList in db by specified user id. Returns id of new list and error
	Create(userId int, list model.TodoList) (int, error)

	// GetAll returns error and slice of model.TodoList by specified user id
	GetAll(userId int) ([]model.TodoList, error)

	// GetById returns error and model of model.TodoList by specified user and list ids
	GetById(userId, listId int) (model.TodoList, error)

	// DeleteById deletes model.TodoList from db by specified list and user ids. Returns an error if there is one.
	// This method also deletes all the items from specified list
	DeleteById(userId, listId int) error

	// Update updates model.TodoList in db by specified list, user ids and model of updated item. Returns an error if there is one
	Update(userId, listId int, list model.UpdateListInput) error
}

type TodoItem interface {
	// Create creates a new model.TodoItem in db by specified list id. Returns id of new item model and error
	Create(listId int, item model.TodoItem) (int, error)

	// GetAll returns error and slice of model.TodoItem by specified list id
	GetAll(listId int) ([]model.TodoItem, error)

	// GetById returns error and model of model.TodoItem by specified list and item ids
	GetById(listId int, itemId int) (model.TodoItem, error)

	// DeleteById deletes model.TodoItem from db by specified list and item ids. Returns an error if there is one
	DeleteById(listId int, itemId int) error

	// Update updates model.TodoItem in db by specified list, item ids and model of updated item. Returns an error if there is one
	Update(listId int, itemId int, item model.UpdateItemInput) error
}

// Repository for working with db. Contains interfaces for splitting business logic
type Repository struct {
	Authorization
	TodoList
	TodoItem
}

// NewRepository returns a pointer on new instance of Repository
func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		TodoList:      postgres.NewTodoListPostgres(db),
		TodoItem:      postgres.NewTodoItemPostgres(db),
	}
}
