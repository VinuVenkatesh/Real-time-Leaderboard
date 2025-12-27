package middleware

import (
	"net/http"

	"github.com/VinuVenkatesh/Real-time-Leaderboard/backend/services"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Authorization")
		url := r.URL.Path
		if url == "/login" || url == "/register" {
			next.ServeHTTP(w, r)
			return
		}
		if jwtToken == "" {
			http.Error(w, "Missing Authorization Header", http.StatusUnauthorized)
			return
		}
		token, err := services.ValidateJWT(jwtToken)
		if err != nil || !token.Valid {
			http.Error(w, "Invalid Token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
