package repository

import "go-todo-app/domain"

type TodoRepo interface {
	GetAll() ([]*domain.Todo, error)
	Create(title string) (*domain.Todo, error)
	Update(id int, status domain.TodoStatus) (*domain.Todo, error)
}
