package domain

import "time"

type UserID uint
type UserEmail string
type UserPassword string

type User struct {
	ID        UserID       `json:"id"`
	Email     UserEmail    `json:"email"`
	Password  UserPassword `json:"password"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
}
