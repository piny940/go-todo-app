package repository

import "go-todo-app/domain"

type ITodoRepo interface {
	List() ([]*domain.Todo, error)
	Create(title domain.TodoTitle, status domain.TodoStatus) (*domain.Todo, error)
	Update(id domain.TodoID, title domain.TodoTitle, status domain.TodoStatus) (*domain.Todo, error)
	FindById(id domain.TodoID) (*domain.Todo, error)
}
