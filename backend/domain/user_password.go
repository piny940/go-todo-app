package domain

import "golang.org/x/crypto/bcrypt"

type UserRawPassword string
type UserHashedPassword string

type UserPassword struct {
	HashedPassword UserHashedPassword
}

func NewUserPassword(rawPassword UserRawPassword) (*UserPassword, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(string(rawPassword)), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedPassword := UserHashedPassword(string(hash))
	return &UserPassword{HashedPassword: hashedPassword}, nil
}

func (p *UserPassword) Check(rawPassword UserRawPassword) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.HashedPassword), []byte(rawPassword))
	return err == nil
}
