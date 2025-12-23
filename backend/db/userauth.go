package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
)

var DB *sql.DB //upper case to make it public/global

func InitDB() error {
	godotenv.Load(".env") //relative to where go is run
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	createTableSQL := `CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	DB = db
	return nil
}

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

func CloseDB() {
	DB.Close()
}

func GetPasswordHash(username string) (string, error) {
	query := `SELECT password_hash FROM registrations WHERE username=$1`
	var passwordHash string
	if err := DB.QueryRow(query, username).Scan(&passwordHash); err != nil {
		return "", err
	}
	return passwordHash, nil
}
