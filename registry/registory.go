package registry

import (
	"go-todo-app/db"
	"go-todo-app/domain/repository"
	"go-todo-app/use_case"
)

type IRegistry interface {
	TodoUseCase() use_case.ITodoUseCase
	TodoRepo() repository.ITodoRepo
}

type registry struct {
	db *db.DB
}

var reg IRegistry

func Init(db *db.DB) {
	reg = &registry{db: db}
}

func GetRegistry() IRegistry {
	return reg
}

func (r *registry) TodoRepo() repository.ITodoRepo {
	return db.NewTodoRepo(r.db)
}

func (r *registry) TodoUseCase() use_case.ITodoUseCase {
	return use_case.NewTodoUseCase(r.TodoRepo())
}
