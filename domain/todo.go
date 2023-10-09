package domain

import "time"

type TodoID int
type TodoTitle string
type TodoStatus int

const (
	Active TodoStatus = iota
	Completed
)

type Todo struct {
	ID        TodoID     `json:"id"`
	Title     TodoTitle  `json:"title"`
	Status    TodoStatus `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func NewTodo(title TodoTitle) *Todo {
	return &Todo{
		Title:  title,
		Status: Active,
	}
}
