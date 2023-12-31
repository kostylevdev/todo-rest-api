package domain

import "errors"

type TodoList struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description" db:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoListCreate struct {
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
}

type TodoItemCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type TodoItem struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" binding:"required" db:"title"`
	Description string `json:"description " db:"description"`
	Done        bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type TodoListUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type TodoItemUpdate struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Done        *bool   `json:"done"`
}

func (list TodoListUpdate) ValidateTodoListUpdate() error {
	if list.Title == nil && list.Description == nil {
		return errors.New("update must have title or description")
	}
	return nil
}

func (item TodoItemUpdate) ValidateTodoItemUpdate() error {
	if item.Title == nil && item.Description == nil && item.Done == nil {
		return errors.New("update must have title or description or done")
	}
	return nil
}
