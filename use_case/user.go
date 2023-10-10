package use_case

import (
	"go-todo-app/domain"
	"go-todo-app/domain/repository"
)

type IUserUseCase interface {
	SignUp(email domain.UserEmail, password domain.UserRawPassword) (*domain.User, error)
}

type userUseCase struct {
	userRepo repository.IUserRepo
}

func NewUserUseCase(userRepo repository.IUserRepo) IUserUseCase {
	return &userUseCase{userRepo: userRepo}
}

func (u *userUseCase) SignUp(email domain.UserEmail, password domain.UserRawPassword) (*domain.User, error) {
	user, err := domain.NewUser(email, password)
	if err != nil {
		return nil, err
	}
	return u.userRepo.Create(user.Email, user.EncryptedPassword)
}
