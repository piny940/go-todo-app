package registry

import (
	"go-todo-app/db"
	"go-todo-app/domain/repository"
	"go-todo-app/use_case"
)

type Registry interface {
	TodoUseCase() use_case.TodoUseCase
	TodoRepo() repository.TodoRepo
}

type registry struct {
	db *db.DB
}

var reg Registry

func Init(db *db.DB) {
	reg = &registry{db: db}
}

func GetRegistry() Registry {
	return reg
}

func (r *registry) TodoRepo() repository.TodoRepo {
	return db.NewTodoRepo(r.db)
}

func (r *registry) TodoUseCase() use_case.TodoUseCase {
	return use_case.NewTodoUseCase(r.TodoRepo())
}
