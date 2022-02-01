package service

import (
	"github.com/ArturGumerov/todoApp"
	"github.com/ArturGumerov/todoApp/pkg/repository"
)

type TodoListservice struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListservice {
	return &TodoListservice{repo: repo}

}

func (s *TodoListservice) Create(userId int, list todoApp.TodoList) (int, error) {
	return s.repo.Create(userId, list)
}

func (s *TodoListservice) GetAll(userId int) ([]todoApp.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListservice) GetById(userId, listId int) (todoApp.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListservice) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}

func (s *TodoListservice) Update(userId, listId int, input todoApp.UpdateListInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return s.repo.Update(userId, listId, input)

}
