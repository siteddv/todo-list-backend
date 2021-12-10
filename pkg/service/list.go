package service

import (
	"todolistBackend/pkg/model"
	"todolistBackend/pkg/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(userId int, list model.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListService) GetAll(userId int) ([]model.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (model.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) DeleteById(userId, listId int) error {
	return s.repo.DeleteById(userId, listId)
}

func (s *TodoListService) Update(userId, listId int, list model.UpdateListInput) error {
	if err := list.Validate(); err != nil {
		return err

	}
	return s.repo.Update(userId, listId, list)
}
