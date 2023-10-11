package repository

import "go-todo-app/domain"

type IUserRepo interface {
	List() ([]*domain.User, error)
	Create(email domain.UserEmail, password domain.UserEncryptedPassword) (*domain.User, error)
}
