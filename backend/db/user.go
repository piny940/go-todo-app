package db

import (
	"go-todo-app/domain"
	"time"
)

type userRepo struct {
	db *DB
}

func NewUserRepo(db *DB) *userRepo {
	return &userRepo{db: db}
}

type UserTable struct {
	ID                uint      `json:"id"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func (u *userRepo) List() ([]*domain.User, error) {
	var users = make([]*domain.User, 0)
	rows, err := u.db.Client.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var userTable UserTable
		if err := rows.Scan(
			&userTable.ID,
			&userTable.Email,
			&userTable.EncryptedPassword,
			&userTable.CreatedAt,
			&userTable.UpdatedAt,
		); err != nil {
			return nil, err
		}
		password := domain.UserPassword{
			HashedPassword: domain.UserHashedPassword(userTable.EncryptedPassword),
		}
		user := domain.User{
			ID:        domain.UserID(userTable.ID),
			Email:     domain.UserEmail(userTable.Email),
			Password:  password,
			CreatedAt: userTable.CreatedAt,
			UpdatedAt: userTable.UpdatedAt,
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *userRepo) Create(email domain.UserEmail, password domain.UserHashedPassword) (*domain.User, error) {
	var user domain.User
	if err := u.db.Client.QueryRow(
		"insert into users (email, password) values ($1, $2) returning *",
		email,
		password,
	).Scan(
		&user.ID,
		&user.Email,
		&user.EncryptedPassword,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, err
	}

	return &user, nil
}
