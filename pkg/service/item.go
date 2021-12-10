package service

import (
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(listId int, item model.TodoItem) (int, error) {
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(listId int) ([]model.TodoItem, error) {
	return s.repo.GetAll(listId)
}

func (s *TodoItemService) GetById(listId, itemId int) (model.TodoItem, error) {
	return s.repo.GetById(listId, itemId)
}

func (s *TodoItemService) DeleteById(listId, itemId int) error {
	return s.repo.DeleteById(listId, itemId)
}

func (s *TodoItemService) Update(listId int, itemId int, item model.UpdateItemInput) error {
	if err := item.Validate(); err != nil {
		return err

	}
	return s.repo.Update(listId, itemId, item)
}
