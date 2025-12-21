package routes

import (
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/handlers"
)

func RegistrationRoutes() {
	http.HandleFunc("/register", handlers.RegistrationHandler)
}
