package model

import "errors"

// TodoList is a model of TodoList
type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

// UsersLists is a model of 3rd table linking User and TodoList
type UsersLists struct {
	Id     int
	UserId int
	ListId int
}

// TodoItem is a model of TodoItem
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// ListsItems is a model of 3rd table linking TodoItem and TodoList
type ListsItems struct {
	Id     int
	ListId int
	ItemId int
}

// UpdateListInput is a partially model of TodoList for update an exciting TodoList
type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// UpdateItemInput is a partially model of TodoItem for update an exciting TodoItem
type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        bool    `json:"done"`
}

// Validate checks correctness of UpdateListInput on nil of any fields and return err if model is invalid
func (input UpdateListInput) Validate() error {
	if input.Title == nil || input.Description == nil {
		return errors.New("update list structure has no values")
	}
	return nil
}

// Validate checks correctness of UpdateItemInput on nil of any fields and return err if model is invalid
func (input UpdateItemInput) Validate() error {
	if input.Title == nil || input.Description == nil {
		return errors.New("update item structure has no values")
	}
	return nil
}
