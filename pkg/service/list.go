package service

import (
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

// TodoListService contains repository for working with items in db
type TodoListService struct {
	repo repository.TodoList
}

// NewTodoListService returns pointer on a new instance of TodoListService
func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

// Create creates a new model.TodoList in db by specified user id. Returns id of new list and error
func (s *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

// GetAll returns error and slice of model.TodoList by specified user id
func (s *TodoListService) GetAll(userId int) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}

// GetById returns error and model of model.TodoList by specified user and list ids
func (s *TodoListService) GetById(userId, listId int) (model.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

// DeleteById deletes model.TodoList from db by specified list and user ids. Returns an error if there is one.
// This method also deletes all the items from specified list
func (s *TodoListService) DeleteById(userId, listId int) error {
	return s.repo.DeleteById(userId, listId)
}

// Update updates model.TodoList in db by specified list, user ids and model of updated item. Returns an error if there is one
func (s *TodoListService) Update(userId, listId int, list model.UpdateListInput) error {
	if err := list.Validate(); err != nil {
		return err

	}
	return s.repo.Update(userId, listId, list)
}
