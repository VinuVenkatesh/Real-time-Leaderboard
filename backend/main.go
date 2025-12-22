package main

import (
	"log"
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/routes"
)

const PORT = ":1234"

func main() {
	db.InitDB()
	routes.RegistrationRoutes()
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal("Listen and Serve: ", err)
	}
	defer db.CloseDB()
}
