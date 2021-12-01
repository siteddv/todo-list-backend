package repository

import (
	"github.com/jmoiron/sqlx"
	todo "todolistBackend"
	postgres "todolistBackend/pkg/repository/postgres"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	DeleteById(userId, listId int) error
	Update(userId, listId int, list todo.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(listId int) ([]todo.TodoItem, error)
	GetById(listId int, itemId int) (todo.TodoItem, error)
	DeleteById(listId int, itemId int) error
	Update(listId int, itemId int, item todo.UpdateItemInput) error
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
