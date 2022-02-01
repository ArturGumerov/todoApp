package repository

import (
	"github.com/ArturGumerov/todoApp"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user todoApp.User) (int, error)
	GetUser(username, password string) (todoApp.User, error)
}
type TodoList interface {
	Create(userId int, list todoApp.TodoList) (int, error)
	GetAll(userId int) ([]todoApp.TodoList, error)
	GetById(userId, listId int) (todoApp.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todoApp.UpdateListInput) error
}
type TodoItem interface {
	Create(listId int, item todoApp.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todoApp.TodoItem, error)
	GetById(userId, itemId int) (todoApp.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input todoApp.UpdateItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
