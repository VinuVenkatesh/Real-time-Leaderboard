package services

import (
	"errors"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(username string, password string) error {
	exists, err := db.CheckUserExists(username)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("username already exists")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	if err := db.Registration(username, string(passwordHash)); err != nil {
		return err
	}
	return nil
}

func AuthenticateUser(username string, password string) (bool, error) {
	storedHash, err := db.GetPasswordHash(username)
	if err != nil {
		return false, errors.New("User does not exist")
	}
	if err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password)); err != nil {
		return false, errors.New("Incorrect password")
	}
	return true, nil
}
