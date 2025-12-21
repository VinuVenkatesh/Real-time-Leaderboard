package main
import (
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/routes"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/db"
)

func main() {
	db.InitDB()
	routes.RegistrationRoutes()

}