package todoListBackend

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

type ListsItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type UpdateItemInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (input UpdateListInput) Validate() error {
	if input.Title == nil || input.Description == nil {
		return errors.New("update list structure has no values")
	}
	return nil
}

func (input UpdateItemInput) Validate() error {
	if input.Title == nil || input.Description == nil {
		return errors.New("update item structure has no values")
	}
	return nil
}
