package main

import (
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/routes"
)

func main() {
	db.InitDB()
	routes.RegistrationRoutes()

}
