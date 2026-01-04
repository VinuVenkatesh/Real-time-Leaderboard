package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
)

var DB *sql.DB //upper case to make it public/global

func InitDB() error {
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
	createTableRegistrationsSQL := `CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableRegistrationsSQL)
	if err != nil {
		panic(err)
	}
	createTableSportsActivitiesSQL := `
		CREATE TABLE IF NOT EXISTS sports_activities (
		id SERIAL PRIMARY KEY,
		activity_name VARCHAR(100) UNIQUE NOT NULL
	);`
	_, err = db.Exec(createTableSportsActivitiesSQL)
	if err != nil {
		panic(err)
	}
	DB = db
	return nil
}

func CloseDB() {
	DB.Close()
}
