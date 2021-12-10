package repository

import (
	"github.com/jmoiron/sqlx"
	"todolistBackend/pkg/model"
	postgres "todolistBackend/pkg/repository/postgres"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId, listId int) (model.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, list model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(listId int) ([]model.TodoItem, error)
	GetById(listId int, itemId int) (model.TodoItem, error)
	DeleteById(listId int, itemId int) error
	Update(listId int, itemId int, item model.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		TodoList:      postgres.NewTodoListPostgres(db),
		TodoItem:      postgres.NewTodoItemPostgres(db),
	}
}
