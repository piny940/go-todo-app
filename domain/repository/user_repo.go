package repository

import "go-todo-app/domain"

type IUserRepo interface {
	Create(email domain.UserEmail, password domain.UserEncryptedPassword) (*domain.User, error)
}
