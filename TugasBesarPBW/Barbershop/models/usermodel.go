package models

import (
	"go-web-native/config"
	"go-web-native/entities"

	"golang.org/x/crypto/bcrypt"
)

func GetUserByUsername(username string) (entities.User, error) {
	var user entities.User
	err := config.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = ?", username).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (entities.User, error) {
	var user entities.User
	err := config.DB.QueryRow("SELECT id, username, email, password FROM users WHERE email = ?", email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByUsernameOrEmail(username, email string) (entities.User, error) {
	var user entities.User
	err := config.DB.QueryRow("SELECT id, username, email, password FROM users WHERE username = ? OR email = ?", username, email).Scan(&user.Id, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(username, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	_, err = config.DB.Exec("INSERT INTO users (username, email, password) VALUES (?, ?, ?)", username, email, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
