package service

import (
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

// TodoItemService contains repository for working with items in db
type TodoItemService struct {
	repo repository.TodoItem
}

// NewTodoItemService returns pointer on a new instance of TodoItemService
func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

// Create creates a new model.TodoItem in db by specified list id. Returns id of new item model and error
func (s *TodoItemService) Create(listId int, item model.TodoItem) (int, error) {
	return s.repo.Create(listId, item)
}

// GetAll returns error and slice of model.TodoItem by specified list id
func (s *TodoItemService) GetAll(listId int) ([]model.TodoItem, error) {
	return s.repo.GetAll(listId)
}

// GetById returns error and model of model.TodoItem by specified list and item ids
func (s *TodoItemService) GetById(listId, itemId int) (model.TodoItem, error) {
	return s.repo.GetById(listId, itemId)
}

// DeleteById deletes model.TodoItem from db by specified list and item ids. Returns an error if there is one
func (s *TodoItemService) DeleteById(listId, itemId int) error {
	return s.repo.DeleteById(listId, itemId)
}

// Update updates model.TodoItem in db by specified list, item ids and model of updated item. Returns an error if there is one
func (s *TodoItemService) Update(listId int, itemId int, item model.UpdateItemInput) error {
	if err := item.Validate(); err != nil {
		return err

	}
	return s.repo.Update(listId, itemId, item)
}
