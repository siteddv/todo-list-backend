package service

import (
	todo "todolistBackend"
	"todolistBackend/pkg/repository"
)

type TodoItemService struct {
	repo repository.TodoItem
}

func NewTodoItemService(repo repository.TodoItem) *TodoItemService {
	return &TodoItemService{repo: repo}
}

func (s *TodoItemService) Create(listId int, item todo.TodoItem) (int, error) {
	return s.repo.Create(listId, item)
}

func (s *TodoItemService) GetAll(listId int) ([]todo.TodoItem, error) {
	return s.repo.GetAll(listId)
}

func (s *TodoItemService) GetById(listId, itemId int) (todo.TodoItem, error) {
	return s.repo.GetById(listId, itemId)
}

func (s *TodoItemService) DeleteById(listId, itemId int) error {
	return s.repo.DeleteById(listId, itemId)
}

func (s *TodoItemService) Update(listId int, itemId int, item todo.UpdateItemInput) error {
	if err := item.Validate(); err != nil {
		return err

	}
	return s.repo.Update(listId, itemId, item)
}
