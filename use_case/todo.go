package use_case

import (
	"go-todo-app/domain"
	"go-todo-app/domain/repository"
)

type ITodoUseCase interface {
	List() ([]*domain.Todo, error)
	Create(title domain.TodoTitle) (*domain.Todo, error)
	// Complete(id domain.TodoID) (*domain.Todo, error)
}

type todoUseCase struct {
	todoRepo repository.ITodoRepo
}

func NewTodoUseCase(todoRepo repository.ITodoRepo) ITodoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

func (t *todoUseCase) List() ([]*domain.Todo, error) {
	return t.todoRepo.List()
}

func (t *todoUseCase) Create(title domain.TodoTitle) (*domain.Todo, error) {
	todo := domain.NewTodo(title)
	return t.todoRepo.Create(todo.Title, todo.Status)
}

// func (t *todoUseCase) Update(id domain.TodoID) (*domain.Todo, error) {
// 	return t.todoRepo.Update(id, status)
// }
