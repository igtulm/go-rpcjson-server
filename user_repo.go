package main

import (
	"time"
)

type IUserRepository interface {
	// Gets user by login
	GetByLogin(login string) (*User, error)
	// Creates a new user with login specified
	Create(login string) (string, error)
	// Changes login to another for login specified
	Update(login string, newLogin string) error
}

type UserRepository struct {
	conn IDbConnector
}

func NewUserRepository(conn IDbConnector) IUserRepository {
	return &UserRepository{conn: conn}
}

func (repo *UserRepository) GetByLogin(login string) (*User, error) {
	rows, err := repo.conn.GetDb().Query(`SELECT id, login, created_at from users WHERE login = $1`, login)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user User
	var createdAt time.Time
	// always be a single result because of unique index constraint
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Login, &createdAt); err != nil {
			return nil, err
		}
		user.CreatedAt = createdAt.Unix()
	}
	return &user, nil
}

func (repo *UserRepository) Create(login string) (string, error) {
	now := time.Now()
	var ID string
	row := repo.conn.GetDb().QueryRow(`INSERT INTO users(login, created_at) VALUES($1, $2) RETURNING id`,
		login,
		now)
	err := row.Scan(&ID)
	if err != nil {
		return "", err
	}
	return ID, nil
}

func (repo *UserRepository) Update(login string, newLogin string) error {
	_, err := repo.conn.GetDb().Exec(`UPDATE users SET login = $1 WHERE login = $2`, newLogin, login)
	if err != nil {
		return err
	}
	return nil
}
