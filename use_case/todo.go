package use_case

import (
	"go-todo-app/domain"
	"go-todo-app/domain/repository"
)

type todoUseCase struct {
	todoRepo repository.TodoRepo
}

func NewTodoUseCase(todoRepo repository.TodoRepo) *todoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

func (t *todoUseCase) Index() ([]*domain.Todo, error) {
	return t.todoRepo.Index()
}

func (t *todoUseCase) Create(title string) (*domain.Todo, error) {
	return t.todoRepo.Create(title)
}

func (t *todoUseCase) Update(id int, status domain.TodoStatus) (*domain.Todo, error) {
	return t.todoRepo.Update(id, status)
}
