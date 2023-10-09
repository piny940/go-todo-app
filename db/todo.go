package db

import (
	"go-todo-app/domain"
	"go-todo-app/domain/repository"
)

type todoRepo struct {
	db *DB
}

func NewTodoRepo(db *DB) repository.ITodoRepo {
	return &todoRepo{
		db: db,
	}
}

func (t *todoRepo) List() ([]*domain.Todo, error) {
	rows, err := t.db.Client.Query("SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	todos := make([]*domain.Todo, 0)
	for rows.Next() {
		var todo domain.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, &todo)
	}
	return todos, nil
}

func (t *todoRepo) Create(title domain.TodoTitle, status domain.TodoStatus) (*domain.Todo, error) {
	var todo domain.Todo
	if err := t.db.Client.QueryRow(
		"INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING *",
		title, status,
	).Scan(
		&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *todoRepo) Update(id domain.TodoID, title domain.TodoTitle, status domain.TodoStatus) (*domain.Todo, error) {
	var todo domain.Todo
	if err := t.db.Client.QueryRow(
		"UPDATE todos SET title = $1, status = $2 WHERE id = $3 RETURNING *",
		title, status, id,
	).Scan(
		&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &todo, nil
}

func (t *todoRepo) FindById(id domain.TodoID) (*domain.Todo, error) {
	var todo domain.Todo
	if err := t.db.Client.QueryRow(
		"SELECT * FROM todos WHERE id = $1", id,
	).Scan(
		&todo.ID, &todo.Title, &todo.Status, &todo.CreatedAt, &todo.UpdatedAt,
	); err != nil {
		return nil, err
	}
	return &todo, nil
}
