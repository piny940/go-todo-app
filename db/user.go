package db

import "go-todo-app/domain"

type userRepo struct {
	db *DB
}

func NewUserRepo(db *DB) *userRepo {
	return &userRepo{db: db}
}

func (u *userRepo) List() ([]*domain.User, error) {
	var users = make([]*domain.User, 0)
	rows, err := u.db.Client.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.EncryptedPassword,
			&user.CreatedAt,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *userRepo) Create(email domain.UserEmail, password domain.UserEncryptedPassword) (*domain.User, error) {
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
