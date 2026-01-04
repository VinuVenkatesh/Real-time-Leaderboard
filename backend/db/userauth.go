package db

import (
	"fmt"

	_ "github.com/lib/pq"
)

func Registration(username string, passwordHash string) error {
	query := `INSERT INTO registrations (username, password_hash) VALUES ($1, $2)`
	_, err := DB.Exec(query, username, passwordHash)
	if err != nil {
		return err
	}
	return nil
}

func CheckUserExists(username string) (bool, error) {
	query := `SELECT COUNT(*) FROM registrations WHERE username=$1`
	var count int
	if err := DB.QueryRow(query, username).Scan(&count); err != nil {
		fmt.Println("Error checking user existence:", err)
		return false, err
	}
	return count > 0, nil
}

func GetPasswordHash(username string) (string, error) {
	query := `SELECT password_hash FROM registrations WHERE username=$1`
	var passwordHash string
	if err := DB.QueryRow(query, username).Scan(&passwordHash); err != nil {
		return "", err
	}
	return passwordHash, nil
}
