package routes

import (
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/handlers"
)

func UserAuthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/register", handlers.RegistrationHandler)
	mux.HandleFunc("/login", handlers.LoginHandler)
}
