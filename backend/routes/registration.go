package routes

import (
	"net/http"
	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/handlers"
)

func registrationRoutes(){
	http.HandleFunc("/register", handlers.RegistrationHandler)
}