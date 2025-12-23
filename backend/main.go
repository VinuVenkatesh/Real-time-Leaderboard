package main

import (
	"log"
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/middleware"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/routes"
)

const PORT = ":1234"

func main() {
	db.InitDB()
	mux := http.NewServeMux()
	handler := middleware.CORS(mux)
	routes.UserAuthRoutes(mux)
	if err := http.ListenAndServe(PORT, handler); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
	defer db.CloseDB()
}
