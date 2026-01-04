package main

import (
	"log"
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/middleware"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/routes"
	"github.com/joho/godotenv"
)

const PORT = ":1234"

func main() {
	godotenv.Load(".env") // Load environment variables from .env file with relative path
	db.InitDB()
	db.InitSportsActivities()
	mux := http.NewServeMux()
	handler := middleware.CORS(middleware.JWTAuth(mux))
	routes.UserAuthRoutes(mux)
	if err := http.ListenAndServe(PORT, handler); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}

	defer db.CloseDB()
}
