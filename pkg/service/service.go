package service

import (
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
