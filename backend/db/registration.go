import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"github.com/joho/godotenv"
)

const (
	host = "localhost"
	port = 5432
)

func InitDB() error {
	godotenv.Load("../.env")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+ "password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	createTableSQL:= `CREATE TABLE IF NOT EXISTS registrations (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) UNIQUE NOT NULL,
		password_hash VARCHAR(255) NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func registration(username string, passwordHash string){
	
}
	

