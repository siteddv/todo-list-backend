package service

import (
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	// CreateUser creates model.User in DB using specified user model. It returns new user id and error
	CreateUser(user model.User) (int, error)

	// GenerateToken generate token for signing in user. Returns a complete token and error
	GenerateToken(username, password string) (string, error)

	// ParseToken decrypts token and returns id of signed in user and error
	ParseToken(token string) (int, error)
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

// Service is a core service containing all the services
type Service struct {
	Authorization
	TodoList
	TodoItem
}

// NewService returns a pointer on a new instance of Service
func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem),
	}
}
