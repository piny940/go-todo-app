package domain

import (
	"go-todo-app/lib"
	"time"
)

type UserID uint
type UserEmail string
type UserRawPassword string
type UserEncryptedPassword string

type User struct {
	ID                UserID                `json:"id"`
	Email             UserEmail             `json:"email"`
	EncryptedPassword UserEncryptedPassword `json:"password"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
}

func NewUser(email UserEmail, rawPassword UserRawPassword) (*User, error) {
	encryptedPassword, err := lib.EncryptPassword(string(rawPassword))
	if err != nil {
		return nil, err
	}
	return &User{
		Email:             email,
		EncryptedPassword: UserEncryptedPassword(encryptedPassword),
	}, nil
}
