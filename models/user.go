package models

import (
	"errors"
	"my-webap/utils" // Ensure this path is correct or update it to the correct module path
)

type User struct {
	ID       int
	Email    string
	Password string
}

var users = make(map[string]User)

func AuthenticateUser(email, password string) (*User, error) {
	user, exists := users[email]
	if !exists {
		return nil, errors.New("user not found")
	}

	if !utils.CheckPasswordHash(password, user.Password) {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func CreateUser(email, password string) error {
	if _, exists := users[email]; exists {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	users[email] = User{
		ID:       len(users) + 1,
		Email:    email,
		Password: hashedPassword,
	}

	return nil
}
